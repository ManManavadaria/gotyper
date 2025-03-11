package main

import (
	"log"

	"github.com/ManManavadaria/gotyper/player"

	"github.com/gdamore/tcell/v2"
)

func main() {
	App := player.NewApplication()
	defer App.TviewApp.Stop()

	App.TviewApp.SetBeforeDrawFunc(func(screen tcell.Screen) bool {
		screen.Clear()
		return false
	})

	if err := App.CreateWelcome(); err != nil {
		log.Fatal(err)
	}

	// // go func() {
	if err := App.TviewApp.Run(); err != nil {
		log.Fatal(err)
	}
}
