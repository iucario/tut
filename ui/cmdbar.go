package ui

import (
	"strings"

	"github.com/RasmusLindroth/tut/config"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type CmdBar struct {
	tutView *TutView
	View    *tview.InputField
}

func NewCmdBar(tv *TutView) *CmdBar {
	c := &CmdBar{
		tutView: tv,
		View:    NewInputField(tv.tut.Config),
	}
	c.View.SetAutocompleteFunc(c.Autocomplete)
	c.View.SetDoneFunc(c.DoneFunc)

	return c
}

func (c *CmdBar) GetInput() string {
	return strings.TrimSpace(c.View.GetText())
}

func (c *CmdBar) ShowError(s string) {
	c.View.SetFieldTextColor(c.tutView.tut.Config.Style.WarningText)
	c.View.SetText(s)
}

func (c *CmdBar) ShowMsg(s string) {
	c.View.SetFieldTextColor(c.tutView.tut.Config.Style.CommandText)
	c.View.SetText(s)
}

func (c *CmdBar) ClearInput() {
	c.View.SetFieldTextColor(c.tutView.tut.Config.Style.CommandText)
	c.View.SetText("")
}

func (c *CmdBar) Back() {
	c.ClearInput()
	c.View.Autocomplete()
	c.tutView.PrevFocus()
}

func (c *CmdBar) DoneFunc(key tcell.Key) {
	if key == tcell.KeyTAB {
		return
	}
	input := c.GetInput()
	parts := strings.Split(input, " ")
	if len(parts) == 0 {
		return
	}
	switch parts[0] {
	case ":q":
		fallthrough
	case ":quit":
		c.tutView.tut.App.Stop()
	case ":compose":
		c.tutView.ComposeCommand()
		c.ClearInput()
		c.View.Autocomplete()
	case ":edit":
		c.ClearInput()
		c.View.Autocomplete()
		c.Back()
		c.tutView.EditCommand()
	case ":blocking":
		c.tutView.BlockingCommand()
		c.Back()
	case ":bookmarks", ":saved":
		c.tutView.BookmarksCommand()
		c.Back()
	case ":favorited":
		c.tutView.FavoritedCommand()
		c.Back()
	case ":boosts":
		c.tutView.BoostsCommand()
		c.Back()
	case ":favorites":
		c.tutView.FavoritesCommand()
		c.Back()
	case ":following":
		c.tutView.FollowingCommand()
		c.Back()
	case ":followers":
		c.tutView.FollowersCommand()
		c.Back()
	case ":history":
		c.tutView.HistoryCommand()
		c.Back()
	case ":newer":
		c.tutView.LoadNewerCommand()
		c.Back()
	case ":clear-notifications":
		c.tutView.ClearNotificationsCommand()
		c.Back()
	case ":close-window":
		c.tutView.CloseWindowCommand()
		c.Back()
	case ":move-window", ":mv":
		if len(parts) < 2 {
			break
		}
		switch parts[1] {
		case "left", "up", "l", "u":
			c.tutView.MoveWindowLeft()
			c.Back()
		case "right", "down", "r", "d":
			c.tutView.MoveWindowRight()
			c.Back()
		case "home", "h":
			c.tutView.MoveWindowHome()
			c.Back()
		case "end", "e":
			c.tutView.MoveWindowEnd()
			c.Back()
		}
	case ":list-placement":
		if len(parts) < 2 {
			break
		}
		switch parts[1] {
		case "top":
			c.tutView.ListPlacementCommand(config.ListPlacementTop)
			c.Back()
		case "right":
			c.tutView.ListPlacementCommand(config.ListPlacementRight)
			c.Back()
		case "bottom":
			c.tutView.ListPlacementCommand(config.ListPlacementBottom)
			c.Back()
		case "left":
			c.tutView.ListPlacementCommand(config.ListPlacementLeft)
			c.Back()
		}
	case ":list-split":
		if len(parts) < 2 {
			break
		}
		switch parts[1] {
		case "column":
			c.tutView.ListSplitCommand(config.ListColumn)
			c.Back()
		case "row":
			c.tutView.ListSplitCommand(config.ListRow)
			c.Back()
		}
	case ":muting":
		c.tutView.MutingCommand()
		c.Back()
	case ":requests":
		c.tutView.FollowRequestsCommand()
		c.Back()
	case ":proportions":
		if len(parts) < 3 {
			break
		}
		c.tutView.ProportionsCommand(parts[1], parts[2])
		c.Back()
	case ":profile":
		c.tutView.ProfileCommand()
		c.Back()
	case ":preferences":
		c.tutView.PreferencesCommand()
		c.ClearInput()
		c.View.Autocomplete()
	case ":timeline", ":tl":
		if len(parts) < 2 {
			break
		}
		switch parts[1] {
		case "local", "l":
			c.tutView.LocalCommand()
			c.Back()
		case "federated", "f":
			c.tutView.FederatedCommand()
			c.Back()
		case "special-all", "sa":
			c.tutView.SpecialCommand(true, true)
			c.Back()
		case "special-boosts", "sb":
			c.tutView.SpecialCommand(true, false)
			c.Back()
		case "special-replies", "sr":
			c.tutView.SpecialCommand(false, true)
			c.Back()
		case "direct", "d":
			c.tutView.DirectCommand()
			c.Back()
		case "home", "h":
			c.tutView.HomeCommand()
			c.Back()
		case "notifications", "n":
			c.tutView.NotificationsCommand()
			c.Back()
		case "mentions", "m":
			c.tutView.MentionsCommand()
			c.Back()
		case "favorited", "fav":
			c.tutView.FavoritedCommand()
			c.Back()
		}
	case ":tag":
		if len(parts) < 2 {
			break
		}
		tParts := strings.TrimSpace(strings.Join(parts[1:], " "))
		if len(tParts) == 0 {
			break
		}
		c.tutView.TagCommand(tParts)
		c.Back()
	case ":tags":
		c.tutView.TagsCommand()
		c.Back()
	case ":window":
		if len(parts) < 2 {
			break
		}
		c.tutView.WindowCommand(parts[1])
		c.Back()
	case ":user":
		if len(parts) < 2 {
			break
		}
		user := strings.TrimSpace(parts[1])
		if len(user) == 0 {
			break
		}
		c.tutView.Timeline.AddFeed(
			NewUserSearchFeed(c.tutView, user),
		)
		c.Back()
	case ":refetch":
		c.tutView.RefetchCommand()
		c.Back()
	case ":stick-to-top":
		c.tutView.ToggleStickToTop()
		c.Back()
	case ":follow-tag":
		if len(parts) < 2 {
			break
		}
		tag := strings.TrimSpace(parts[1])
		if len(tag) == 0 {
			break
		}
		c.tutView.TagFollowCommand(parts[1])
		c.Back()
	case ":unfollow-tag":
		if len(parts) < 2 {
			break
		}
		tag := strings.TrimSpace(parts[1])
		if len(tag) == 0 {
			break
		}
		c.tutView.TagUnfollowCommand(parts[1])
		c.Back()
	case ":lists":
		c.tutView.ListsCommand()
		c.Back()
	case ":help", ":h":
		c.tutView.PageFocus = c.tutView.PrevPageFocus
		c.tutView.SetPage(HelpFocus)
		c.ClearInput()
		c.View.Autocomplete()
	}
}

func (c *CmdBar) Autocomplete(curr string) []string {
	var entries []string
	words := strings.Split(":blocking,:boosts,:bookmarks,:clear-notifications,:compose,:favorites,:favorited,:follow-tag,:followers,:following,:help,:h,:history,:move-window,:lists,:list-placement,:list-split,:muting,:newer,:preferences,:profile,:proportions,:refetch,:requests,:saved,:stick-to-top,:tag,:timeline,:tl,:unfollow-tag,:user,:window,:quit,:q", ",")
	if curr == "" {
		return entries
	}

	if len(curr) > 2 && curr[:3] == ":tl" {
		words = strings.Split(":tl home,:tl notifications,:tl local,:tl federated,:tl direct,:tl mentions,:tl favorited,:tl special-all,:tl special-boosts,:tl-special-replies", ",")
	}
	if len(curr) > 8 && curr[:9] == ":timeline" {
		words = strings.Split(":timeline home,:timeline notifications,:timeline local,:timeline federated,:timeline direct,:timeline mentions,:timeline favorited,:timeline special-all,:timeline special-boosts,:timeline special-replies", ",")
	}
	if len(curr) > 14 && curr[:15] == ":list-placement" {
		words = strings.Split(":list-placement top,:list-placement right,:list-placement bottom,:list-placement left", ",")
	}
	if len(curr) > 10 && curr[:11] == ":list-split" {
		words = strings.Split(":list-split row,:list-split column", ",")
	}

	if len(curr) > 11 && curr[:12] == ":move-window" {
		words = strings.Split(":move-window left,:move-window right,:move-window up,:move-window down,:move-window home,:move-window end", ",")
	}
	if len(curr) > 2 && curr[:3] == ":mv" {
		words = strings.Split(":mv left,:mv right,:mv up,:mv down,:mv home,:mv end", ",")
	}

	for _, word := range words {
		if strings.HasPrefix(strings.ToLower(word), strings.ToLower(curr)) {
			entries = append(entries, word)
		}
	}
	if len(entries) < 1 {
		entries = nil
	}
	return entries
}
