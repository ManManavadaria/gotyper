package player

import (
	"github.com/ManManavadaria/gotyper/utils"
	"github.com/ManManavadaria/gotyper/widgets"

	"github.com/navidys/tvxwidgets"
	"github.com/rivo/tview"
)

func CreateWelcome(app *tview.Application, gauge *tvxwidgets.ActivityModeGauge) error {

	const welcomeSign = `welcome...`

	signWi := tview.NewTextView().SetText(welcomeSign)
	menuWi := tview.NewList().
		AddItem("Let's start", "", 0, func() {
			CreateSinglePlayer(app, gauge)
		}).
		AddItem("exit", "exit the app", 0, func() {
			app.Stop()
		})

	signW, signH := utils.StringDimensions(welcomeSign)
	menuW, menuH := 32, menuWi.GetItemCount()*2
	welcomeLayout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tview.NewBox(), 0, 1, false).
		AddItem(Center(signW, signH, signWi), 0, 1, false).
		AddItem(Center(menuW, menuH, menuWi), 0, 1, true).
		AddItem(tview.NewBox(), 0, 1, false)

	widgets.CreateActivityProgressBar(gauge, app, "Initiating the application...")

	pages := tview.NewPages().
		AddPage("root1", welcomeLayout, true, true).
		AddPage("root2", gauge, true, false)

	app.SetRoot(pages, true)
	return nil
}

// func ExitApp(app *tview.Application) error {
// 	fmt.Println("exitinf application=====")
// 	app.Stop()
// 	return nil
// }
