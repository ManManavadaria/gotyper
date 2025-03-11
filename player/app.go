package player

import (
	"github.com/navidys/tvxwidgets"
	"github.com/rivo/tview"
)

type App struct {
	TviewApp *tview.Application
	Gauge    *tvxwidgets.ActivityModeGauge
	Pages    *tview.Pages
	Layouts  map[string]*tview.Flex
}

func NewApplication() *App {
	return &App{
		TviewApp: tview.NewApplication(),
		Gauge:    tvxwidgets.NewActivityModeGauge(),
		Pages:    nil,
		Layouts:  make(map[string]*tview.Flex),
	}
}
