package layout

import (
	"ani-track/src/services"
	"ani-track/src/tui/ui/widgets/actionSection"
	"ani-track/src/tui/ui/widgets/actionSection/addPage"
	"ani-track/src/tui/ui/widgets/actionSection/downloadPage"
	"ani-track/src/tui/ui/widgets/actionSection/modifyPage"
	"ani-track/src/tui/ui/widgets/captionSection"
	"ani-track/src/tui/ui/widgets/historylistSection"
	"ani-track/src/tui/ui/widgets/identifierSection"
	"ani-track/src/tui/ui/widgets/overviewSection"
	"ani-track/src/tui/ui/widgets/watchlistSection"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	ROWS    = []int{1, 1, 1, 0, 11, 1}
	COLUMNS = []int{4, 0, 15, 15, 20, 20, 15, 20, 50}
)

func MasterLayout(app *tview.Application) *tview.Grid {
	masterHeader := captionSection.Caption("[cyan::b]═══ Welcome to Ani-Track! ═══[white]  - Your Anime, Manga & Novel Tracker")
	detailHeader := captionSection.Caption("[yellow::b]Detailed View[-::-]")

	statusBar := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter).
		SetText("[yellow]Tab:[white] Switch Pane  │  [yellow]a:[white] Add  │  [yellow]m/Enter:[white] Modify  │  [yellow]d:[white] Export  │  [yellow]q:[white] Quit  │  [yellow]?:[white] Help")

	masterLayout := tview.NewGrid().
		SetRows(ROWS...).
		SetColumns(COLUMNS...).
		SetBorders(true)

	overview := overviewSection.Overview()
	historyTable := historylistSection.HistoryList()
	watchlistTable := watchlistSection.WatchList()

	var currentSelected *services.Entry

	getSelected := func() *services.Entry {
		return currentSelected
	}

	refreshData := func() {
		historyEntries, _ := services.GetHistory()
		watchlistEntries, _ := services.GetWatchlist()

		historyTable.UpdateData(historyEntries)
		watchlistTable.UpdateData(watchlistEntries)

		if currentSelected != nil {
			if updated, err := services.GetEntry(currentSelected.Index); err == nil {
				currentSelected = updated
			} else {
				if len(historyEntries) > 0 {
					currentSelected = &historyEntries[0]
				} else if len(watchlistEntries) > 0 {
					currentSelected = &watchlistEntries[0]
				} else {
					currentSelected = nil
				}
			}
		} else {
			if len(historyEntries) > 0 {
				currentSelected = &historyEntries[0]
			} else if len(watchlistEntries) > 0 {
				currentSelected = &watchlistEntries[0]
			}
		}

		overview.SetEntry(currentSelected)
	}

	actionMenu := actionSection.ActionSection(app, masterLayout, getSelected, refreshData)

	historyTable.SetSelectionChangedFunc(func(row, column int) {
		if e := historyTable.GetSelectedEntry(); e != nil {
			currentSelected = e
			overview.SetEntry(e)
		}
	})

	watchlistTable.SetSelectionChangedFunc(func(row, column int) {
		if e := watchlistTable.GetSelectedEntry(); e != nil {
			currentSelected = e
			overview.SetEntry(e)
		}
	})

	historyTable.SetSelectedFunc(func(row, column int) {
		if currentSelected != nil {
			modifyPage.Menu(app, masterLayout, getSelected, refreshData)
		}
	})

	watchlistTable.SetSelectedFunc(func(row, column int) {
		if currentSelected != nil {
			modifyPage.Menu(app, masterLayout, getSelected, refreshData)
		}
	})

	refreshData()
	if e := historyTable.GetSelectedEntry(); e != nil {
		currentSelected = e
		overview.SetEntry(e)
	}

	identifierSection.Identifier(masterLayout)

	masterLayout.
		AddItem(masterHeader, 0, 0, 1, 9, 0, 0, false).
		AddItem(detailHeader, 2, 8, 1, 1, 0, 0, false).
		AddItem(historyTable.Table, 3, 0, 1, 8, 0, 0, true).
		AddItem(overview.Grid, 3, 8, 1, 1, 0, 0, false).
		AddItem(watchlistTable.Table, 4, 0, 1, 8, 0, 0, false).
		AddItem(actionMenu.Grid, 4, 8, 1, 1, 0, 0, false).
		AddItem(statusBar, 5, 0, 1, 9, 0, 0, false)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		currentFocus := app.GetFocus()
		if event.Key() == tcell.KeyTab {
			if currentFocus == historyTable.Table {
				app.SetFocus(watchlistTable.Table)
				if e := watchlistTable.GetSelectedEntry(); e != nil {
					currentSelected = e
					overview.SetEntry(e)
				}
				return nil
			} else if currentFocus == watchlistTable.Table {
				app.SetFocus(actionMenu.BtnAdd)
				return nil
			} else {
				app.SetFocus(historyTable.Table)
				if e := historyTable.GetSelectedEntry(); e != nil {
					currentSelected = e
					overview.SetEntry(e)
				}
				return nil
			}
		}

		if event.Key() == tcell.KeyBacktab {
			if currentFocus == historyTable.Table {
				app.SetFocus(actionMenu.BtnExit)
				return nil
			} else if currentFocus == watchlistTable.Table {
				app.SetFocus(historyTable.Table)
				if e := historyTable.GetSelectedEntry(); e != nil {
					currentSelected = e
					overview.SetEntry(e)
				}
				return nil
			} else {
				app.SetFocus(watchlistTable.Table)
				if e := watchlistTable.GetSelectedEntry(); e != nil {
					currentSelected = e
					overview.SetEntry(e)
				}
				return nil
			}
		}

		switch event.Rune() {
		case 'a', 'A':
			addPage.Menu(app, masterLayout, refreshData)
			return nil
		case 'm', 'M':
			modifyPage.Menu(app, masterLayout, getSelected, refreshData)
			return nil
		case 'd', 'D':
			downloadPage.Menu(app, masterLayout)
			return nil
		case 'q', 'Q':
			app.Stop()
			return nil
		case '?':
			showHelpModal(app, masterLayout)
			return nil
		}

		if event.Key() == tcell.KeyF1 {
			showHelpModal(app, masterLayout)
			return nil
		}

		return event
	})

	return masterLayout
}

func showHelpModal(app *tview.Application, layout tview.Primitive) {
	helpText := `[yellow::b]Ani-Track Keyboard Shortcuts[-::-]

  [yellow]Tab / Shift-Tab[-]: Cycle focus between History, Watchlist, Actions
  [yellow]Up / Down[-]:       Navigate items in tables & buttons
  [yellow]Enter[-]:           Modify selected entry / Trigger button
  [yellow]a[-]:               Add new entry (Anime, Manga, Novel)
  [yellow]m[-]:               Modify or Delete selected entry
  [yellow]d[-]:               Export / Download collection to text file
  [yellow]q[-]:               Quit application
  [yellow]? / F1[-]:          Show this help dialog
  [yellow]Esc[-]:             Cancel / Close current modal dialog`

	modal := tview.NewModal().
		SetText(helpText).
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			app.SetRoot(layout, true)
		})

	app.SetRoot(modal, true)
}

func PlaceHolder(placeholder string) *tview.TextView {
	return tview.NewTextView().SetText(placeholder).SetTextAlign(tview.AlignCenter)
}
