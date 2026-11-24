package downloadPage

import (
	"ani-track/src/services"
	"ani-track/src/tui/ui/widgets/actionSection/addPage"
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Menu(app *tview.Application, layout tview.Primitive) {
	header := tview.NewTextView().
		SetDynamicColors(true).
		SetText("Export your entire anime and manga collection to a formatted text file with ASCII banner and tables.")

	location := tview.NewInputField().
		SetLabel("Export Path: ").
		SetText("~/Downloads/Ani-Track.txt").
		SetFieldWidth(45).
		SetFieldBackgroundColor(tcell.ColorDefault).
		SetFieldTextColor(tcell.ColorWhite)

	statusView := tview.NewTextView().
		SetDynamicColors(true)

	form := tview.NewForm()
	form.SetBorder(true).
		SetTitle(" [::b]Export Collection (Esc to Cancel) ").
		SetTitleAlign(tview.AlignCenter).
		SetBorderColor(tcell.ColorDarkCyan)

	form.
		AddFormItem(header).
		AddFormItem(location).
		AddFormItem(statusView).
		AddButton("Download", func() {
			target := location.GetText()
			err := services.Convert(target)
			if err != nil {
				statusView.SetTextColor(tcell.ColorRed).
					SetText(fmt.Sprintf(" Export error: %v", err))
				return
			}

			successModal := tview.NewModal().
				SetText(fmt.Sprintf("List successfully exported to:\n%s", target)).
				AddButtons([]string{"OK"}).
				SetDoneFunc(func(buttonIndex int, buttonLabel string) {
					app.SetRoot(layout, true)
				})
			app.SetRoot(successModal, true)
		}).
		AddButton("Cancel", func() {
			app.SetRoot(layout, true)
		})

	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			app.SetRoot(layout, true)
			return nil
		}
		return event
	})

	modal := addPage.ModalGrid(form, 65, 14)
	app.SetRoot(modal, true)
}
