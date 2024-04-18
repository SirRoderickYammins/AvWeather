package core

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Global Variables for Gobrief
var (
	App        *GoBrief
	AppVersion = "GoBrief v0.1"
)

// Main Application struct: TView is the tview libray's own construct.
// PageHolder is the page that will be rendered at any given time.
type GoBrief struct {
	TView      *tview.Application
	PageHolder *tview.Pages
	PageName   string
	// Add Pilot Profiles
}

// Set the default UI color theme.
// Doing this was poorly documented...
func defaultTheme() {
	tview.Styles.PrimitiveBackgroundColor = tcell.Color16
	tview.Styles.BorderColor = tcell.ColorGrey
	tview.Styles.ContrastBackgroundColor = tcell.ColorBlue
	tview.Styles.MoreContrastBackgroundColor = tcell.Color16
}

// Initialize the main application
func (g *GoBrief) Init() {
	g.TView.SetRoot(g.PageHolder, true).EnableMouse(true).SetFocus(g.PageHolder)
	defaultTheme()
}
