package main

import (
	"log"

	"github.com/ManManavadaria/gotyper/player"
	"github.com/ManManavadaria/gotyper/widgets"

	"github.com/gdamore/tcell/v2"
	"github.com/navidys/tvxwidgets"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	defer app.Stop()

	gauge := tvxwidgets.NewActivityModeGauge()

	widgets.CreateActivityProgressBar(gauge, app, "Initializing the applicatio...")

	app.SetBeforeDrawFunc(func(screen tcell.Screen) bool {
		screen.Clear()
		return false
	})

	if err := player.CreateWelcome(app, gauge); err != nil {
		log.Fatal(err)
	}

	// // go func() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
