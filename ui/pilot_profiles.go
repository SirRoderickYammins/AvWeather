package ui

// TODO: Implement Pilot profile data structure

import (
	"github.com/gdamore/tcell/v2"
	"github.com/xpndrobserved/gobrief/core"
)

func ShowPilotProfilePage() {
	pilotPage := core.CreatePilotProfilePage()

	core.App.TView.SetFocus(pilotPage.Grid)
	core.App.PageHolder.AddAndSwitchToPage("PilotProfiles", pilotPage.Grid, true)
	core.App.PageName = "pilotProfiles"

	core.App.TView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if core.App.PageName == "pilotProfiles" {
			if event.Rune() == 'q' {
				ShowMainPage()
			}
		}
		return event
	})
}
