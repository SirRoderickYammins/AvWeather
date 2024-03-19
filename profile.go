package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	help = `Press [green]Ctrl+N[-] to create a new pilot profile.`
)

func Profile(nextSlide func()) (title string, content tview.Primitive) {
	pilotListRoot := tview.NewTreeNode("Pilots").
		SetSelectable(false).
		SetColor(tcell.ColorFireBrick)

	tree := tview.NewTreeView().
		SetRoot(pilotListRoot).
		SetCurrentNode(pilotListRoot)

	frame := tview.NewFrame(tview.NewBox()).
		AddText(help, true, tview.AlignLeft, tcell.ColorWhite)

	pilotForm := tview.NewForm().
		AddInputField("First Name", "", 20, nil, nil).
		AddInputField("Last Name", "", 20, nil, nil).
		AddDropDown("Home Airport", []string{"KBWI", "KFME", "KIAD", "KADW"}, 0, nil).
		AddButton("Save", nil).
		AddButton("Quit", func() {
			app.Stop()
		})

	flex := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(tree, 0, 1, true).
		AddItem(frame, 0, 1, false).
		AddItem(pilotForm, 0, 1, true)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlR {
			flex.AddItem(pilotForm, 0, 1, true)
			app.SetFocus(pilotForm)
			return nil
		}
		return event
	})

	return "Profiles", flex
}
