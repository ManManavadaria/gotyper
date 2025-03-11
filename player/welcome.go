package player

import (
	"github.com/ManManavadaria/gotyper/utils"

	"github.com/rivo/tview"
)

func (a *App) CreateWelcome() error {

	const welcomeSign = `welcome...`

	signWi := tview.NewTextView().SetText(welcomeSign)
	menuWi := tview.NewList().
		AddItem("Let's start", "", 0, func() {
			a.CreateSinglePlayer()
		}).
		AddItem("exit", "exit the app", 0, func() {
			a.TviewApp.Stop()
		})

	signW, signH := utils.StringDimensions(welcomeSign)
	menuW, menuH := 32, menuWi.GetItemCount()*2
	a.Layouts["welcome"] = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tview.NewBox(), 0, 1, false).
		AddItem(Center(signW, signH, signWi), 0, 1, false).
		AddItem(Center(menuW, menuH, menuWi), 0, 1, true).
		AddItem(tview.NewBox(), 0, 1, false)

	a.CreateActivityProgressBar("Initiating the application...")

	a.Pages = tview.NewPages().
		AddPage("welcome", a.Layouts["welcome"], true, false).
		AddPage("loader", a.Gauge, false, true)

	a.TviewApp.SetRoot(a.Pages, true)
	return nil
}
