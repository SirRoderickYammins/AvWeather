package main

import (
	"github.com/rivo/tview"
)

type Pilot struct {
	FirstName   string
	LastName    string
	HomeAirport string
}

var pilots []Pilot

func newPilotForm() (content tview.Primitive) {
	form := tview.NewForm().
		AddInputField("First Name", "firstname", 20, nil, nil).
		AddInputField("Last Name", "lastname", 20, nil, nil).
		AddDropDown("Home Airport", []string{"KBWI", "KFME", "KIAD", "KADW"}, 0, nil)

	addPilot := func() {
		selected := form.GetFormItemByLabel("Home Airport").GetLabel()

		thisPilot := Pilot{
			FirstName:   form.GetFormItemByLabel("firstname").GetLabel(),
			LastName:    form.GetFormItemByLabel("lastname").GetLabel(),
			HomeAirport: selected,
		}
		pilots = append(pilots, thisPilot)
	}

	form.AddButton("Add Pilot", addPilot)

	form.SetBorder(true).SetTitle("Add Pilot")
	return form
}
