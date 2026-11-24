package watchlistSection

import (
	"ani-track/src/services"
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type WatchlistTable struct {
	*tview.Table
	entries []services.Entry
}

func WatchList() *WatchlistTable {
	table := tview.NewTable().
		SetBorders(false).
		SetSelectable(true, false).
		SetSeparator(tview.Borders.Vertical)

	table.SetBorder(true).
		SetTitle(" [::b]Watch List (Ongoing / Planned) ").
		SetTitleAlign(tview.AlignLeft).
		SetBorderColor(tcell.ColorDarkCyan)

	table.SetSelectedStyle(tcell.StyleDefault.
		Background(tcell.ColorDarkCyan).
		Foreground(tcell.ColorWhite).
		Bold(true))

	wt := &WatchlistTable{
		Table: table,
	}

	return wt
}

func (wt *WatchlistTable) UpdateData(entries []services.Entry) {
	wt.entries = entries
	wt.Clear()

	headers := []string{"ID", "Title", "Type", "Rating", "Started", "Finished", "Status", "Progress", "Notes"}
	for col, h := range headers {
		cell := tview.NewTableCell(" " + h + " ").
			SetTextColor(tcell.ColorYellow).
			SetSelectable(false).
			SetAlign(tview.AlignCenter)

		if col == 1 {
			cell.SetAlign(tview.AlignLeft).SetExpansion(1)
		} else if col == 8 {
			cell.SetAlign(tview.AlignLeft).SetMaxWidth(30)
		}
		wt.SetCell(0, col, cell)
	}
	wt.SetFixed(1, 0)

	for rowIdx, e := range entries {
		row := rowIdx + 1

		idCell := tview.NewTableCell(fmt.Sprintf("%2d", e.Index)).
			SetAlign(tview.AlignCenter).
			SetTextColor(tcell.ColorGray)

		titleCell := tview.NewTableCell(" " + e.Title).
			SetAlign(tview.AlignLeft).
			SetTextColor(tcell.ColorWhite).
			SetExpansion(1)

		typeCell := tview.NewTableCell(e.Category).
			SetAlign(tview.AlignCenter).
			SetTextColor(tcell.ColorDarkTurquoise)

		ratingCell := tview.NewTableCell(e.Rating).
			SetAlign(tview.AlignCenter).
			SetTextColor(tcell.ColorGreen)

		startedCell := tview.NewTableCell(e.Started).
			SetAlign(tview.AlignCenter).
			SetTextColor(tcell.ColorLightGray)

		finishedCell := tview.NewTableCell(e.Finished).
			SetAlign(tview.AlignCenter).
			SetTextColor(tcell.ColorLightGray)

		statusCell := tview.NewTableCell(e.Status).
			SetAlign(tview.AlignCenter).
			SetTextColor(tcell.ColorAqua)

		progressCell := tview.NewTableCell(e.Progress).
			SetAlign(tview.AlignCenter).
			SetTextColor(tcell.ColorPurple)

		notesCell := tview.NewTableCell(e.Notes).
			SetAlign(tview.AlignLeft).
			SetMaxWidth(30).
			SetTextColor(tcell.ColorLightGray)

		wt.SetCell(row, 0, idCell)
		wt.SetCell(row, 1, titleCell)
		wt.SetCell(row, 2, typeCell)
		wt.SetCell(row, 3, ratingCell)
		wt.SetCell(row, 4, startedCell)
		wt.SetCell(row, 5, finishedCell)
		wt.SetCell(row, 6, statusCell)
		wt.SetCell(row, 7, progressCell)
		wt.SetCell(row, 8, notesCell)
	}
}

func (wt *WatchlistTable) GetSelectedEntry() *services.Entry {
	row, _ := wt.GetSelection()
	if row > 0 && row <= len(wt.entries) {
		return &wt.entries[row-1]
	}
	return nil
}
