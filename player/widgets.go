package player

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
	"golang.org/x/term"
)

func (a *App) CreateActivityProgressBar(msg string) error {

	termW, termH, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		log.Println(err)
		return err
	}
	x := termW/2 - 50/2
	y := termH/2 - 3/2

	a.Gauge.SetTitle(msg)
	a.Gauge.SetPgBgColor(tcell.ColorOrange)
	// a.Gauge.SetRect(termW/2, termH/2, 50, 3)
	a.Gauge.SetRect(x, y, 50, 3)
	// a.Gauge.SetRect(10, 4, 50, 3)

	a.Gauge.SetBorder(true)

	update := func() {
		ctx, cancle := context.WithTimeout(context.Background(), time.Millisecond*400)
		defer cancle()
		tick := time.NewTicker(10 * time.Millisecond)
		for {
			select {
			case <-tick.C:
				a.Gauge.Pulse()
				a.TviewApp.Draw()
			case <-ctx.Done():
				tick.Stop()
				a.Pages.SwitchToPage("welcome")
				a.TviewApp.QueueUpdateDraw(func() {
					a.Pages.RemovePage("loader")
				})
				a.TviewApp.SetFocus(a.Layouts["welcome"])
				return
			}
		}
	}
	go update()
	return nil
}
