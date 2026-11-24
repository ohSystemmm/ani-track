package actionSection

import (
	"ani-track/src/services"
	"ani-track/src/tui/ui/widgets/actionSection/addPage"
	"ani-track/src/tui/ui/widgets/actionSection/downloadPage"
	"ani-track/src/tui/ui/widgets/actionSection/modifyPage"

	"github.com/rivo/tview"
)

func AddButton(app *tview.Application, layout tview.Primitive, onRefresh func()) *tview.Button {
	return tview.NewButton("Add").
		SetSelectedFunc(func() {
			addPage.Menu(app, layout, onRefresh)
		})
}

func ModifyButton(app *tview.Application, layout tview.Primitive, getSelected func() *services.Entry, onRefresh func()) *tview.Button {
	return tview.NewButton("Modify").
		SetSelectedFunc(func() {
			modifyPage.Menu(app, layout, getSelected, onRefresh)
		})
}

func DownloadButton(app *tview.Application, layout tview.Primitive) *tview.Button {
	return tview.NewButton("Download").
		SetSelectedFunc(func() {
			downloadPage.Menu(app, layout)
		})
}

func ExitButton(app *tview.Application) *tview.Button {
	return tview.NewButton("Exit").
		SetSelectedFunc(func() {
			app.Stop()
		})
}
