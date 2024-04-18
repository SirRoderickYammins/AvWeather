package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/xpndrobserved/gobrief/core"
)

type SettingsPage struct {
	Grid *tview.Grid
}

func createSettingsPage() *SettingsPage {
	grid := tview.NewGrid().
		SetColumns(0)

	settingsListRoot := tview.NewTreeNode("Settings")

	tree := tview.NewTreeView().
		SetRoot(settingsListRoot).
		SetCurrentNode(settingsListRoot)

	grid.AddItem(tree, 0, 0, 1, 1, 0, 100, true).
		SetBorders(true)

	settingsPage := SettingsPage{grid}

	return &settingsPage
}

func ShowSettingsPage() {
	settingsPage := createSettingsPage()

	core.App.TView.SetFocus(settingsPage.Grid)
	core.App.PageHolder.AddAndSwitchToPage("Settings", settingsPage.Grid, true)
	core.App.PageName = "settings"

	core.App.TView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if core.App.PageName == "settings" {
			if event.Rune() == 'q' {
				ShowMainPage()
			}
		}
		return event
	})
}
