package identifierSection

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func base(object *tview.TextView, text string) *tview.TextView {
	return object.
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter).
		SetTextColor(tcell.ColorYellow).
		SetText(text)
}

func labelID() *tview.TextView {
	return base(tview.NewTextView(), "ID")
}
func labelName() *tview.TextView {
	return base(tview.NewTextView(), "Title")
}
func labelType() *tview.TextView {
	return base(tview.NewTextView(), "Category")
}
func labelRating() *tview.TextView {
	return base(tview.NewTextView(), "Rating")
}
func labelStarted() *tview.TextView {
	return base(tview.NewTextView(), "Started")
}
func labelFinished() *tview.TextView {
	return base(tview.NewTextView(), "Finished")
}
func labelStatus() *tview.TextView {
	return base(tview.NewTextView(), "Status")
}
func labelNotes() *tview.TextView {
	return base(tview.NewTextView(), "Notes")
}
