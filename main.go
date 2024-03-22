package main

import (
	"log"

	"github.com/rivo/tview"
	"github.com/xpndrobserved/gobrief/core"
	"github.com/xpndrobserved/gobrief/ui"
)

func main() {
	// Initialize Main Application
	core.App = &core.GoBrief{
		TView:      tview.NewApplication(),
		PageHolder: tview.NewPages(),
	}

	ui.ShowSplashScreen()

	core.App.Init()

	if err := core.App.TView.Run(); err != nil {
		log.Println(err)
	}
}
