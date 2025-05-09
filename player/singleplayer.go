package player

import (
	"fmt"
	"strings"
	"time"

	"github.com/ManManavadaria/gotyper/utils"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (a *App) CreateSinglePlayer() error {
	text, err := GenerateText()
	if err != nil {
		return err
	}

	state := NewState(text)

	statsWis := [...]*tview.TextView{
		tview.NewTextView().SetText("wpm: 0"),
		tview.NewTextView().SetText("Time: 0s"),
	}

	pages := tview.NewPages().
		AddPage("modal", tview.NewModal().
			SetText("Play again?").
			SetBackgroundColor(tcell.ColorDefault).
			AddButtons([]string{"yes", "exit"}).
			SetDoneFunc(func(index int, label string) {
				switch index {
				case 0:
					utils.Check(a.CreateSinglePlayer())
				case 1:
					utils.Check(a.CreateWelcome())
				}
			}), false, false)

	var textWis []*tview.TextView
	for _, word := range state.Words {
		// textWis = append(textWis, tview.NewTextView().SetText(word).SetDynamicColors(true))
		textWis = append(textWis, tview.NewTextView().SetText(toNarrowFullWidth(word)).SetWordWrap(true).SetDynamicColors(true))
	}

	currInput := ""
	inputWi := tview.NewInputField().
		SetFieldBackgroundColor(tcell.ColorDefault)
	inputWi.
		SetChangedFunc(func(text string) {
			if state.StartTime.IsZero() {
				state.Start()
				go func() {
					ticker := time.NewTicker(100 * time.Millisecond)
					for range ticker.C {
						if state.CurrWord == len(state.Words) {
							ticker.Stop()
							return
						}
						a.TviewApp.QueueUpdateDraw(func() {
							statsWis[0].SetText(fmt.Sprintf("wpm: %.0f", state.Wpm()))
							statsWis[1].SetText(fmt.Sprintf("time: %.02fs", time.Since(state.StartTime).Seconds()))
						})
					}
				}()
			}

			if len(currInput) < len(text) {
				if len(text) > len(state.Words[state.CurrWord]) || state.Words[state.CurrWord][len(text)-1] != text[len(text)-1] {
					state.IncError()
				}
			}

			i := state.CurrWord
			diff := paintDiff(state.Words[i], text)
			go func(i int, diff string) {
				a.TviewApp.QueueUpdateDraw(func() {
					textWis[i].SetText(diff).SetWordWrap(false)
				})
			}(i, diff)

			if text == state.Words[state.CurrWord] {
				if state.CurrWord == len(state.Words)-1 {
					state.End()
					pages.ShowPage("modal")
				} else {
					state.NextWord()
					inputWi.SetText("")
				}
			}

			currInput = text
		})

	layout := tview.NewFlex()
	statsFrame := tview.NewFlex().SetDirection(tview.FlexRow)
	statsFrame.SetBorder(true).SetBorderPadding(1, 1, 1, 1).SetTitle("STATS")
	for _, statsWi := range statsWis {
		statsFrame.AddItem(statsWi, 1, 1, false)
	}
	layout.AddItem(statsFrame, 0, 1, false)

	secondColumn := tview.NewFlex().SetDirection(tview.FlexColumn)
	textsLayout := tview.NewFlex()
	for _, textWi := range textWis {
		textsLayout.AddItem(textWi.SetWrap(true), len(textWi.GetText(true)), 1, false)
	}
	textsLayout.SetBorder(true)

	secondColumn.AddItem(textsLayout, 0, 3, false)
	inputWi.SetBorder(true)
	secondColumn.AddItem(inputWi, 0, 0, true)
	layout.AddItem(secondColumn, 0, 3, true)

	pages.AddPage("game", layout, true, true).SendToBack("game")
	a.TviewApp.SetRoot(pages, true)

	keybindings(a.TviewApp, a.CreateWelcome)
	return nil
}

// paintDiff returns an tview-painted string displaying the difference
func paintDiff(toColor string, differ string) string {
	var h string
	h = ":"
	var colorText string

	for i := range differ {
		if i >= len(toColor) || differ[i] != toColor[i] {
			colorText += "[" + h + "red]"
		} else {
			colorText += "[" + h + "green]"
		}

		colorText += toNarrowFullWidth(string(differ[i]))
	}

	colorText += "[-:-:-]"

	if len(differ) < len(toColor) {
		colorText += toNarrowFullWidth(toColor[len(differ):])
	}

	return colorText
}

func toFullWidth(text string) string {
	var result strings.Builder
	for _, r := range text {
		if r >= 33 && r <= 126 { // ASCII printable characters
			result.WriteRune(r + 0xFEE0) // Convert to full-width
		} else {
			result.WriteRune(r) // Keep non-ASCII characters unchanged
		}
	}
	return result.String()
}

func toNarrowFullWidth(text string) string {
	return text
}

// func toNarrowFullWidth(text string) string {
// 	var result strings.Builder
// 	text = strings.TrimSpace(text)
// 	for _, r := range text {
// 		if r >= 33 && r <= 126 { // Convert only ASCII printable characters
// 			result.WriteRune(r + 0xFEE0) // Convert to full-width
// 			result.WriteRune('\u200B')   // Insert zero-width space to reduce spacing
// 		} else {
// 			result.WriteRune(r) // Keep other characters unchanged
// 		}
// 	}
// 	// log.Println("bigtext->", result.String(), "--->")
// 	return result.String()
// }
