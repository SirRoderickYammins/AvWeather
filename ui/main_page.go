package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/xpndrobserved/gobrief/core"
)

// MainPage Struct containing a tview grid for flex boxes
type MainPage struct {
	Grid *tview.Grid
}

var sampleText = "hello"

// Set up the main page.
func createMainPage() *MainPage {
	grid := tview.NewGrid()

	menuItemList := tview.NewList().
		AddItem("Briefings", "Get Your Personalized Weather Briefing", rune(42), showModal).
		AddItem("IDK", "", rune(42), nil)

	homeWeatherFrame := tview.NewFrame(tview.NewBox().
		SetBorder(true).
		SetTitle(sampleText + " Weather and Forecast").
		SetTitleColor(tcell.ColorBlue))

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(menuItemList, 0, 1, true).
		AddItem(homeWeatherFrame, 0, 1, false)

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
