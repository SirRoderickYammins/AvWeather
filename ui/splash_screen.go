package ui

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/xpndrobserved/gobrief/core"
)

type SplashScreen struct {
	Flex *tview.Flex
}

const (
	logo = `
 ________  ________  ________  ________  ___  _______   ________ 
|\   ____\|\   __  \|\   __  \|\   __  \|\  \|\  ___ \ |\  _____\
\ \  \___|\ \  \|\  \ \  \|\ /\ \  \|\  \ \  \ \   __/|\ \  \__/ 
 \ \  \  __\ \  \\\  \ \   __  \ \   _  _\ \  \ \  \_|/_\ \   __\
  \ \  \|\  \ \  \\\  \ \  \|\  \ \  \\  \\ \  \ \  \_|\ \ \  \_|
   \ \_______\ \_______\ \_______\ \__\\ _\\ \__\ \_______\ \__\ 
    \|_______|\|_______|\|_______|\|__|\|__|\|__|\|_______|\|__|
`
	message = "The Definitive Pilot's Interface for Weather Briefing"
	cont    = `[orange]Press Enter to Continue`
)

// Function to create the splash screen
func createSplashScreen() *SplashScreen {
	lines := strings.Split(logo, "\n")
	logoWidth := 0
	logoHeight := len(lines)
	for _, line := range lines {
		if len(line) > logoWidth {
			logoWidth = len(line)
		}
	}
	logoBox := tview.NewTextView().
		SetTextColor(tcell.ColorBlue)

	fmt.Fprint(logoBox, logo)

	// Create a frame for the subtitle and navigation infos.
	frame := tview.NewFrame(tview.NewBox()).
		SetBorders(0, 0, 0, 0, 0, 0).
		AddText("", true, tview.AlignCenter, tcell.ColorWhite).
		AddText(core.AppVersion, true, tview.AlignCenter, tcell.ColorGreen).
		AddText("", true, tview.AlignCenter, tcell.ColorWhite).
		AddText(message, true, tview.AlignCenter, tcell.ColorWhite).
		AddText("", true, tview.AlignCenter, tcell.ColorWhite).
		AddText(cont, true, tview.AlignCenter, tcell.ColorWhite)

	// Create a Flex layout that centers the logo and subtitle.
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tview.NewBox(), 0, 7, false).
		AddItem(tview.NewFlex().
			AddItem(tview.NewBox(), 0, 1, false).
			AddItem(logoBox, logoWidth, 1, true).
			AddItem(tview.NewBox(), 0, 1, false), logoHeight, 1, true).
		AddItem(frame, 0, 10, false)

		// Instantiate the SplashScreen struct for page rendering.
	splashScreen := SplashScreen{
		Flex: flex,
	}

	return &splashScreen
}

// Initial function to show splashscreen on startup.

func ShowSplashScreen() {
	splashScreen := createSplashScreen()

	core.App.TView.SetFocus(splashScreen.Flex)
	core.App.PageHolder.AddAndSwitchToPage("splashscreen", splashScreen.Flex, true)
	core.App.PageName = "splashScreen"

	// Allow the user to proceed from the splashscreen to the main menu
	core.App.TView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if core.App.PageName == "splashScreen" {
			if event.Key() == tcell.KeyEnter {
				ShowMainPage()
			}
		}

		return event
	})
}
