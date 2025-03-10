package widgets

import (
	"context"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/navidys/tvxwidgets"
	"github.com/rivo/tview"
	"golang.org/x/term"
)

func CreateActivityProgressBar(gauge *tvxwidgets.ActivityModeGauge, app *tview.Application, msg string) *tvxwidgets.ActivityModeGauge {

	termW, termH, _ := term.GetSize(int(os.Stdout.Fd()))

	x := termW/2 - 50/2
	y := termH/2 - 3/2

	gauge.SetTitle(msg)
	gauge.SetPgBgColor(tcell.ColorOrange)
	// gauge.SetRect(termW/2, termH/2, 50, 3)
	gauge.SetRect(x, y, 50, 3)

	gauge.SetBorder(true)

	update := func() {
		ctx, _ := context.WithTimeout(context.Background(), time.Second*2)
		tick := time.NewTicker(20 * time.Millisecond)
		for {
			select {
			case <-tick.C:
				gauge.Pulse()
				app.Draw()
			case <-ctx.Done():
				tick.Stop()
				//NOTE: Implement the pages witch logic

			}
		}
	}
	go update()

	return gauge

}
