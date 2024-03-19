package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Profile(nextSlide func()) (title string, content tview.Primitive) {
	pilotListRoot := tview.NewTreeNode("Pilots").
		SetSelectable(false).
		SetColor(tcell.ColorFireBrick)

	tree := tview.NewTreeView().
		SetRoot(pilotListRoot).
		SetCurrentNode(pilotListRoot)

	tree.SetBorder(true)

	pilotForm := newPilotForm()

	flex := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(tree, 1, 1, true).
		AddItem(pilotForm, 0, 1, false)

	return "Profiles", flex
}
