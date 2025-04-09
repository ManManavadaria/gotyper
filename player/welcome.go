package player

import (
	"github.com/ManManavadaria/gotyper/utils"
	"github.com/gdamore/tcell/v2"

	"github.com/rivo/tview"
)

func (a *App) CreateWelcome() error {

	welcomeSign := `
 ██████╗  ██████╗     ████████╗██╗   ██╗██████╗ ███████╗██████╗ 
██╔════╝ ██╔═══██╗    ╚══██╔══╝╚██╗ ██╔╝██╔══██╗██╔════╝██╔══██╗
██║  ███╗██║   ██║       ██║    ╚████╔╝ ██████╔╝█████╗  ██████╔╝
██║   ██║██║   ██║       ██║     ╚██╔╝  ██╔═══╝ ██╔══╝  ██╔══██╗
╚██████╔╝╚██████╔╝       ██║      ██║   ██║     ███████╗██║  ██║
 ╚═════╝  ╚═════╝        ╚═╝      ╚═╝   ╚═╝     ╚══════╝╚═╝  ╚═╝
                                                                
`

	// 	startSign := `
	// ┏┓┏┳┓┏┓┳┓┏┳┓  ┏┳┓┓┏┏┓┳┳┓┏┓
	// ┗┓ ┃ ┣┫┣┫ ┃    ┃ ┗┫┃┃┃┃┃┃┓
	// ┗┛ ┻ ┛┗┛┗ ┻    ┻ ┗┛┣┛┻┛┗┗┛
	// `
	signWi := tview.NewTextView().SetText(welcomeSign).SetTextAlign(tview.AlignCenter).SetTextColor(tcell.ColorYellow)
	menuWi := tview.NewList().
		// AddItem(startSign, "", 0, func() {
		// 	a.CreateSinglePlayer()
		// }).
		AddItem("[::b]【ＬＥＴＳ　ＳＴＡＲＴ】[::-]", "", 0, func() {
			a.CreateSinglePlayer()
		}).
		AddItem("[::b]【ＱＵＩＴ】[::-]", "", 0, func() {
			a.TviewApp.Stop()
		})
		// ShowSecondaryText(false).SetBackgroundColor(tcell.ColorWheat)

	signW, signH := utils.StringDimensions(welcomeSign)
	menuW, menuH := 32, menuWi.GetItemCount()*2

	menuWi.
		SetMainTextColor(tcell.ColorBlack).
		SetSecondaryTextColor(tcell.ColorWhite).
		ShowSecondaryText(false).
		SetSelectedBackgroundColor(tcell.ColorYellow)

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
