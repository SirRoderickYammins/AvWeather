package main

import (
	"fmt"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Slide func(nextSlide func()) (title string, content tview.Primitive)

const version = `Version [green]1.0`

var app = tview.NewApplication()

var layout = tview.NewFlex()

func main() {
	// ui stylepageSettings
	tview.Styles.PrimitiveBackgroundColor = tcell.Color16
	tview.Styles.ContrastBackgroundColor = tcell.ColorRoyalBlue
	tview.Styles.MoreContrastBackgroundColor = tcell.ColorOrange
	tview.Styles.BorderColor = tcell.ColorRoyalBlue
	tview.Styles.TitleColor = tcell.ColorRoyalBlue
	tview.Styles.GraphicsColor = tcell.ColorRoyalBlue
	tview.Styles.PrimaryTextColor = tcell.ColorWhite
	tview.Styles.SecondaryTextColor = tcell.ColorOrange
	tview.Styles.TertiaryTextColor = tcell.ColorHotPink
	tview.Styles.InverseTextColor = tcell.ColorLemonChiffon
	tview.Styles.ContrastSecondaryTextColor = tcell.ColorPeachPuff
	// The presentation slides.
	slides := []Slide{
		Cover,
		Dashboard,
		Profile,
	}

	pages := tview.NewPages()

	// The bottom row has some info on where we are.
	info := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(false).
		SetHighlightedFunc(func(added, removed, remaining []string) {
			pages.SwitchToPage(added[0])
		})

	// Create the pages for all slides.
	previousSlide := func() {
		slide, _ := strconv.Atoi(info.GetHighlights()[0])
		slide = (slide - 1 + len(slides)) % len(slides)
		info.Highlight(strconv.Itoa(slide)).
			ScrollToHighlight()
	}
	nextSlide := func() {
		slide, _ := strconv.Atoi(info.GetHighlights()[0])
		slide = (slide + 1) % len(slides)
		info.Highlight(strconv.Itoa(slide)).
			ScrollToHighlight()
	}
	for index, slide := range slides {
		title, primitive := slide(nextSlide)
		pages.AddPage(strconv.Itoa(index), primitive, true, index == 0)
		fmt.Fprintf(info, `%d ["%d"][blue]%s[magenta][""]  `, index+1, index, title)
	}
	info.Highlight("0")

	// Create the main layout.
	layout.
		SetDirection(tview.FlexRow).
		AddItem(pages, 0, 1, true).
		AddItem(info, 1, 1, false)

	// Shortcuts to navigate the slides.
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlN {
			nextSlide()
			return nil
		} else if event.Key() == tcell.KeyCtrlP {
			previousSlide()
			return nil
		}
		return event
	})

	layout.SetBackgroundColor(tcell.Color16)

	// Start the application.
	if err := app.SetRoot(layout, true).EnableMouse(true).EnablePaste(true).Run(); err != nil {
		panic(err)
	}
}
