package addPage

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func inputField(label string, placeholder string, width int) *tview.InputField {
	return tview.NewInputField().
		SetLabel(label).
		SetPlaceholder(placeholder).
		SetPlaceholderTextColor(tcell.ColorDarkGray).
		SetFieldWidth(width).
		SetFieldBackgroundColor(tcell.ColorDefault).
		SetFieldTextColor(tcell.ColorWhite)
}

func dropDown(label string, options []string, defaultIdx int) (*tview.DropDown, *string) {
	selected := options[0]
	if defaultIdx >= 0 && defaultIdx < len(options) {
		selected = options[defaultIdx]
	}

	selectedPtr := &selected
	dd := tview.NewDropDown().
		SetLabel(label).
		SetOptions(options, func(text string, index int) {
			*selectedPtr = text
		}).
		SetCurrentOption(defaultIdx).
		SetFieldBackgroundColor(tcell.ColorDefault).
		SetFieldTextColor(tcell.ColorWhite)

	return dd, selectedPtr
}

func textArea(label string, placeholder string) *tview.TextArea {
	ta := tview.NewTextArea().
		SetLabel(label).
		SetPlaceholder(placeholder).
		SetSize(4, 45)
	ta.SetBorder(false)
	return ta
}

func ModalGrid(content tview.Primitive, width, height int) *tview.Grid {
	return tview.NewGrid().
		SetColumns(0, width, 0).
		SetRows(0, height, 0).
		AddItem(content, 1, 1, 1, 1, 0, 0, true)
}
