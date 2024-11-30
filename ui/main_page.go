package ui

import (
	"github.com/rivo/tview"
	"github.com/xpndrobserved/gobrief/core"
)

// MainPage Struct containing a tview grid for flex boxes
type MainPage struct {
	Grid *tview.Grid
}

// Set up the main page.
func createMainPage() *MainPage {
	grid := tview.NewGrid().
		SetColumns(0, 0).
		SetBorders(true)

	menuItemList := tview.NewList().
		AddItem("Briefings", "Get Your Personalized Weather Briefing", rune(42), showModal).
		AddItem("Pilot Profiles", "", rune(42), ShowPilotProfilePage).
		AddItem("Settings", "", rune(42), ShowSettingsPage)

	grid.AddItem(menuItemList, 0, 0, 1, 1, 0, 100, true).
		SetTitle("Dashboard")

	mainPage := MainPage{
		Grid: grid,
	}

	return &mainPage
}

func ShowMainPage() {
	mainPage := createMainPage()
	core.App.TView.SetFocus(mainPage.Grid)
	core.App.PageHolder.AddAndSwitchToPage("MainPage", mainPage.Grid, true)
	core.App.PageName = "mainPage"
}
