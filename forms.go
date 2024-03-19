package main

import (
	"github.com/rivo/tview"
)

func newPilotForm() (content tview.Primitive) {
	form := tview.NewForm().
		AddInputField("First Name", "", 20, nil, nil).
		AddInputField("Last Name", "", 20, nil, nil).
		AddDropDown("Home Airport", []string{"KBWI, KFME, KIAD, KADW"}, 0, nil).
		AddButton("Save", nil).
		AddButton("Quit", func() {
			app.Stop()
		})

	return form
}
