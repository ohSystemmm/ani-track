package overviewSection

import (
	"ani-track/src/services"
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type OverviewView struct {
	*tview.Grid
	titleVal    *tview.TextView
	typeVal     *tview.TextView
	ratingVal   *tview.TextView
	startedVal  *tview.TextView
	finishedVal *tview.TextView
	statusVal   *tview.TextView
	progressVal *tview.TextView
	notesVal    *tview.TextView
}

func Overview() *OverviewView {
	ov := &OverviewView{
		Grid:        tview.NewGrid(),
		titleVal:    createValueView("No item selected"),
		typeVal:     createValueView("-"),
		ratingVal:   createValueView("-"),
		startedVal:  createValueView("-"),
		finishedVal: createValueView("-"),
		statusVal:   createValueView("-"),
		progressVal: createValueView("-"),
		notesVal:    createValueView("-"),
	}

	ov.SetRows(1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0).
		SetColumns(12, 0).
		SetBorders(false)

	row := 0
	ov.addProperty(&row, "Title:", ov.titleVal)
	ov.addProperty(&row, "Type:", ov.typeVal)
	ov.addProperty(&row, "Rating:", ov.ratingVal)
	ov.addProperty(&row, "Started:", ov.startedVal)
	ov.addProperty(&row, "Finished:", ov.finishedVal)
	ov.addProperty(&row, "Status:", ov.statusVal)
	ov.addProperty(&row, "Progress:", ov.progressVal)

	notesLabel := createLabelView("Notes:")
	ov.AddItem(notesLabel, row, 0, 1, 1, 0, 0, false)
	ov.notesVal.SetWrap(true).SetWordWrap(true)
	ov.AddItem(ov.notesVal, row, 1, 1, 1, 0, 0, false)

	return ov
}

func (ov *OverviewView) addProperty(row *int, label string, valView *tview.TextView) {
	ov.AddItem(createLabelView(label), *row, 0, 1, 1, 0, 0, false)
	ov.AddItem(valView, *row, 1, 1, 1, 0, 0, false)
	*row++
}

func createLabelView(text string) *tview.TextView {
	return tview.NewTextView().
		SetDynamicColors(true).
		SetTextColor(tcell.ColorYellow).
		SetText(text)
}

func createValueView(text string) *tview.TextView {
	return tview.NewTextView().
		SetDynamicColors(true).
		SetTextColor(tcell.ColorWhite).
		SetText(text)
}

func (ov *OverviewView) SetEntry(entry *services.Entry) {
	if entry == nil {
		ov.titleVal.SetText("[gray]No item selected[-]")
		ov.typeVal.SetText("[gray]-[-]")
		ov.ratingVal.SetText("[gray]-[-]")
		ov.startedVal.SetText("[gray]-[-]")
		ov.finishedVal.SetText("[gray]-[-]")
		ov.statusVal.SetText("[gray]-[-]")
		ov.progressVal.SetText("[gray]-[-]")
		ov.notesVal.SetText("[gray]-[-]")
		return
	}

	ov.titleVal.SetText(fmt.Sprintf("[yellow::b]%s[-::-]", entry.Title))
	ov.typeVal.SetText(fmt.Sprintf("[cyan]%s[-]", entry.Category))
	ov.ratingVal.SetText(fmt.Sprintf("[green::b]%s[-::-]", entry.Rating))
	ov.startedVal.SetText(entry.Started)
	ov.finishedVal.SetText(entry.Finished)

	statusColor := "green"
	if !entry.IsFinished() {
		statusColor = "aqua"
	}
	ov.statusVal.SetText(fmt.Sprintf("[%s::b]%s[-::-]", statusColor, entry.Status))

	progress := entry.Progress
	if progress == "" {
		progress = "N/A"
	}
	ov.progressVal.SetText(fmt.Sprintf("[purple]%s[-]", progress))

	notes := entry.Notes
	if notes == "" {
		notes = "[gray]No notes recorded.[-]"
	}
	ov.notesVal.SetText(notes)
}
