package ui

import (
	"github.com/rivo/tview"
	"github.com/xpndrobserved/gobrief/core"
)

// MainPage Struct containing a tview grid for flex boxes
type MainPage struct {
	Grid *tview.Grid
}

// Set up the main page. Return pointer to MainPage to keep
// struct's lifetime beyond function scope
func createMainPage() *MainPage {
	grid := tview.NewGrid().
		SetRows(3, 0, 3).
		SetColumns(30, 0, 30)

	menuItemList := tview.NewList().
		AddItem("Briefings", "Get Your Personalized Weather Briefing", rune(42), nil).
		AddItem("IDK", "", rune(42), nil)

	flex := tview.NewFlex().
		AddItem(menuItemList, 0, 1, true)

	grid.AddItem(flex, 0, 0, 5, 5, 0, 0, true).
		SetTitle("Dashboard").
		SetBorder(true)

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
