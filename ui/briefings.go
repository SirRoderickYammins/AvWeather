package ui

import (
	"github.com/rivo/tview"
	"github.com/xpndrobserved/gobrief/core"
)

func errorHandler(buttonIndex int, _ string) {
	if buttonIndex == 0 {
		ShowMainPage()
	}
	if buttonIndex == 1 {
		ShowPilotProfilePage()
	}
}

func showModal() {
	buttons := []string{"Return to Main Menu", "Make a Profile"}

	newModal := tview.NewModal().
		SetText("Unable to generate a personalized briefing.\n\nYou have not set up a pilot profile.").
		AddButtons(buttons).
		SetDoneFunc(errorHandler)
	core.App.TView.SetFocus(newModal)
	core.App.PageHolder.AddAndSwitchToPage("Error", newModal, true)
}
