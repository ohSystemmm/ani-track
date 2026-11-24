package identifierSection

import (
	"github.com/rivo/tview"
)

func Identifier(layout *tview.Grid) {
	layout.
		AddItem(labelID(), 2, 0, 1, 1, 0, 0, false).
		AddItem(labelName(), 2, 1, 1, 1, 0, 0, false).
		AddItem(labelType(), 2, 2, 1, 1, 0, 0, false).
		AddItem(labelRating(), 2, 3, 1, 1, 0, 0, false).
		AddItem(labelStarted(), 2, 4, 1, 1, 0, 0, false).
		AddItem(labelFinished(), 2, 5, 1, 1, 0, 0, false).
		AddItem(labelStatus(), 2, 6, 1, 1, 0, 0, false).
		AddItem(labelNotes(), 2, 7, 1, 1, 0, 0, false)
}
