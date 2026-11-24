package overviewSection

import (
	"github.com/rivo/tview"
)

func textview(text string) *tview.TextView {
	return tview.NewTextView().
		SetDynamicColors(true).
		SetText(text)
}
