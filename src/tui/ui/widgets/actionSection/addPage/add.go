package addPage

import (
	"ani-track/src/services"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Menu(app *tview.Application, layout tview.Primitive, onRefresh func()) {
	categories := []string{"Anime", "Manga", "Novel", "Other"}
	statuses := []string{"Finished", "In Watch", "In Read", "Plan to Watch", "Dropped", "Other"}

	title := inputField("Title:     ", "e.g. Frieren: Beyond Journey's End", 40)
	catDD, selectedCategory := dropDown("Category:  ", categories, 0)
	rating := inputField("Rating:    ", "e.g. 95/100 or 95", 15)
	started := inputField("Started:   ", "e.g. Sep 2023", 20)
	finished := inputField("Finished:  ", "e.g. Mar 2024 or Ongoing", 20)
	statusDD, selectedStatus := dropDown("Status:    ", statuses, 0)
	progress := inputField("Progress:  ", "e.g. Ep 28/28 or Ch 160", 20)
	notes := textArea("Notes:     ", "Add personal review, thoughts, or rating notes...")

	errView := tview.NewTextView().
		SetDynamicColors(true).
		SetTextColor(tcell.ColorRed)

	form := tview.NewForm()
	form.SetBorder(true).
		SetTitle(" [::b]Add New Entry (Esc to Cancel) ").
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
		AddButton("Save", func() {
			titleText := strings.TrimSpace(title.GetText())
			if titleText == "" {
				errView.SetText(" Title is required!")
				return
			}

			entry := services.Entry{
				Title:    titleText,
				Category: *selectedCategory,
				Rating:   services.FormatRating(rating.GetText()),
				Started:  strings.TrimSpace(started.GetText()),
				Finished: strings.TrimSpace(finished.GetText()),
				Status:   *selectedStatus,
				Progress: strings.TrimSpace(progress.GetText()),
				Notes:    strings.TrimSpace(notes.GetText()),
			}

			if err := services.AddEntry(entry); err != nil {
				errView.SetText(" Error saving: " + err.Error())
				return
			}

			if onRefresh != nil {
				onRefresh()
			}
			app.SetRoot(layout, true)
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

	modal := ModalGrid(form, 70, 24)
	app.SetRoot(modal, true)
}
