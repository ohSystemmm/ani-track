package tui

import (
	"ani-track/src/tui/layout"

	"github.com/rivo/tview"
)

func SetupUI() *tview.Application {
	app := tview.NewApplication()
	app.EnableMouse(true)
	masterGrid := layout.MasterLayout(app)
	app.SetRoot(masterGrid, true)
	return app
}
