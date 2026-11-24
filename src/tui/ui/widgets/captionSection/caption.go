package captionSection

import (
	"github.com/rivo/tview"
)

func Caption(text string) *tview.TextView {
	return tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter).
		SetText(text)
}
