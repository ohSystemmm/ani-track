package actionSection

import (
	"ani-track/src/services"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ActionMenu struct {
	*tview.Grid
	BtnAdd      *tview.Button
	BtnModify   *tview.Button
	BtnDownload *tview.Button
	BtnExit     *tview.Button
}

func ActionSection(
	app *tview.Application,
	layout tview.Primitive,
	getSelected func() *services.Entry,
	onRefresh func(),
) *ActionMenu {
	title := tview.NewTextView().
		SetDynamicColors(true).
		SetText("[yellow::b]Actions[-::-]").
		SetTextAlign(tview.AlignCenter)

	btnAdd := AddButton(app, layout, onRefresh)
	btnModify := ModifyButton(app, layout, getSelected, onRefresh)
	btnDownload := DownloadButton(app, layout)
	btnExit := ExitButton(app)

	grid := tview.NewGrid().
		SetRows(1, 1, 1, 3, 1, 3, 0).
		SetColumns(2, 0, 2, 0, 2)

	grid.
		AddItem(title, 1, 1, 1, 3, 0, 0, false).
		AddItem(btnAdd, 3, 1, 1, 1, 0, 0, true).
		AddItem(btnModify, 3, 3, 1, 1, 0, 0, false).
		AddItem(btnDownload, 5, 1, 1, 1, 0, 0, false).
		AddItem(btnExit, 5, 3, 1, 1, 0, 0, false)

	menu := &ActionMenu{
		Grid:        grid,
		BtnAdd:      btnAdd,
		BtnModify:   btnModify,
		BtnDownload: btnDownload,
		BtnExit:     btnExit,
	}

	grid.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		current := app.GetFocus()
		switch event.Key() {
		case tcell.KeyUp:
			if current == btnDownload {
				app.SetFocus(btnAdd)
				return nil
			} else if current == btnExit {
				app.SetFocus(btnModify)
				return nil
			}
		case tcell.KeyDown:
			if current == btnAdd {
				app.SetFocus(btnDownload)
				return nil
			} else if current == btnModify {
				app.SetFocus(btnExit)
				return nil
			}
		case tcell.KeyRight:
			if current == btnAdd {
				app.SetFocus(btnModify)
				return nil
			} else if current == btnDownload {
				app.SetFocus(btnExit)
				return nil
			}
		case tcell.KeyLeft:
			if current == btnModify {
				app.SetFocus(btnAdd)
				return nil
			} else if current == btnExit {
				app.SetFocus(btnDownload)
				return nil
			}
		}
		return event
	})

	return menu
}
