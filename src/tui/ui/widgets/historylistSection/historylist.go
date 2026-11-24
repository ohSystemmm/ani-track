package historylistSection

import (
	"ani-track/src/services"
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type HistoryTable struct {
	*tview.Table
	entries []services.Entry
}

func HistoryList() *HistoryTable {
	table := tview.NewTable().
		SetBorders(false).
		SetSelectable(true, false).
		SetSeparator(tview.Borders.Vertical)

	table.SetBorder(true).
		SetTitle(" [::b]Watch History (Finished) ").
		SetTitleAlign(tview.AlignLeft).
		SetBorderColor(tcell.ColorDarkCyan)

	table.SetSelectedStyle(tcell.StyleDefault.
		Background(tcell.ColorDarkCyan).
		Foreground(tcell.ColorWhite).
		Bold(true))

	ht := &HistoryTable{
		Table: table,
	}

	return ht
}

func (ht *HistoryTable) UpdateData(entries []services.Entry) {
	ht.entries = entries
	ht.Clear()

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
		ht.SetCell(0, col, cell)
	}
	ht.SetFixed(1, 0)

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
			SetTextColor(tcell.ColorGreen)

		progressCell := tview.NewTableCell(e.Progress).
			SetAlign(tview.AlignCenter).
			SetTextColor(tcell.ColorPurple)

		notesCell := tview.NewTableCell(e.Notes).
			SetAlign(tview.AlignLeft).
			SetMaxWidth(30).
			SetTextColor(tcell.ColorLightGray)

		ht.SetCell(row, 0, idCell)
		ht.SetCell(row, 1, titleCell)
		ht.SetCell(row, 2, typeCell)
		ht.SetCell(row, 3, ratingCell)
		ht.SetCell(row, 4, startedCell)
		ht.SetCell(row, 5, finishedCell)
		ht.SetCell(row, 6, statusCell)
		ht.SetCell(row, 7, progressCell)
		ht.SetCell(row, 8, notesCell)
	}
}

func (ht *HistoryTable) GetSelectedEntry() *services.Entry {
	row, _ := ht.GetSelection()
	if row > 0 && row <= len(ht.entries) {
		return &ht.entries[row-1]
	}
	return nil
}
