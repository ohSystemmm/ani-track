package modifyPage

import (
	"ani-track/src/services"
	"ani-track/src/tui/ui/widgets/actionSection/addPage"
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Menu(app *tview.Application, layout tview.Primitive, getSelected func() *services.Entry, onRefresh func()) {
	entry := getSelected()
	if entry == nil {
		modal := tview.NewModal().
			SetText("No entry selected!\nPlease select an item in Watch History or Watch List first.").
			AddButtons([]string{"OK"}).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				app.SetRoot(layout, true)
			})
		app.SetRoot(modal, true)
		return
	}

	categories := []string{"Anime", "Manga", "Novel", "Other"}
	statuses := []string{"Finished", "In Watch", "In Read", "Plan to Watch", "Dropped", "Other"}

	catIdx := 0
	for i, c := range categories {
		if strings.EqualFold(c, entry.Category) {
			catIdx = i
			break
		}
	}

	statusIdx := 0
	for i, s := range statuses {
		if strings.EqualFold(s, entry.Status) {
			statusIdx = i
			break
		}
	}

	title := inputField("Title:     ", entry.Title, 40)
	catDD, selectedCategory := dropDown("Category:  ", categories, catIdx)
	rating := inputField("Rating:    ", entry.Rating, 15)
	started := inputField("Started:   ", entry.Started, 20)
	finished := inputField("Finished:  ", entry.Finished, 20)
	statusDD, selectedStatus := dropDown("Status:    ", statuses, statusIdx)
	progress := inputField("Progress:  ", entry.Progress, 20)
	notes := textArea("Notes:     ", entry.Notes)

	errView := tview.NewTextView().
		SetDynamicColors(true).
		SetTextColor(tcell.ColorRed)

	form := tview.NewForm()
	form.SetBorder(true).
		SetTitle(fmt.Sprintf(" [::b]Modify Entry #%d: %s (Esc to Cancel) ", entry.Index, entry.Title)).
		SetTitleAlign(tview.AlignCenter).
		SetBorderColor(tcell.ColorDarkCyan)

	form.
		AddFormItem(title).
		AddFormItem(catDD).
		AddFormItem(rating).
		AddFormItem(started).
		AddFormItem(finished).
		AddFormItem(statusDD).
		AddFormItem(progress).
		AddFormItem(notes).
		AddFormItem(errView).
		AddButton("Save Changes", func() {
			titleText := strings.TrimSpace(title.GetText())
			if titleText == "" {
				errView.SetText(" Title is required!")
				return
			}

			updated := services.Entry{
				Index:    entry.Index,
				Title:    titleText,
				Category: *selectedCategory,
				Rating:   services.FormatRating(rating.GetText()),
				Started:  strings.TrimSpace(started.GetText()),
				Finished: strings.TrimSpace(finished.GetText()),
				Status:   *selectedStatus,
				Progress: strings.TrimSpace(progress.GetText()),
				Notes:    strings.TrimSpace(notes.GetText()),
			}

			if err := services.UpdateEntry(entry.Index, updated); err != nil {
				errView.SetText(" Error updating: " + err.Error())
				return
			}

			if onRefresh != nil {
				onRefresh()
			}
			app.SetRoot(layout, true)
		}).
		AddButton("Delete", func() {
			confirmModal := tview.NewModal().
				SetText(fmt.Sprintf("Are you sure you want to delete\n\"%s\"?", entry.Title)).
				AddButtons([]string{"Yes, Delete", "Cancel"}).
				SetDoneFunc(func(buttonIndex int, buttonLabel string) {
					if buttonLabel == "Yes, Delete" {
						if err := services.DeleteEntry(entry.Index); err != nil {
							errView.SetText(" Error deleting: " + err.Error())
							return
						}
						if onRefresh != nil {
							onRefresh()
						}
					}
					app.SetRoot(layout, true)
				})
			app.SetRoot(confirmModal, true)
		}).
		AddButton("Cancel", func() {
			app.SetRoot(layout, true)
		})

	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			app.SetRoot(layout, true)
			return nil
		}
		return event
	})

	modal := addPage.ModalGrid(form, 70, 24)
	app.SetRoot(modal, true)
}

func inputField(label, text string, width int) *tview.InputField {
	return tview.NewInputField().
		SetLabel(label).
		SetText(text).
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

func textArea(label, text string) *tview.TextArea {
	ta := tview.NewTextArea().
		SetLabel(label).
		SetText(text, false).
		SetSize(4, 45)
	ta.SetBorder(false)
	return ta
}
