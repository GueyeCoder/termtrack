package ui

import (
	"github.com/GueyeCoder/termtrack/controller"
	"github.com/GueyeCoder/termtrack/model"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func SetupUI(app *tview.Application, taskCtrl *controller.TaskController) tview.Primitive {
	taskList := tview.NewList().
		ShowSecondaryText(false)

	refresh := func() {
		taskList.Clear()
		tasks := taskCtrl.GetAll()
		for _, t := range tasks {
			title := t.Title
			if t.Done {
				title = "[green::b][✓] " + title
			}
			taskList.AddItem(title, "", 0, nil)
		}
	}

	refresh()

	tasks := taskCtrl.GetAll()
	for _, t := range tasks {
		title := t.Title
		if t.Done {
			title = "[green::b][✓] " + title
		}
		taskList.AddItem(title, "", 0, nil)
	}

	footer := tview.NewTextView().
		SetText("[a] Ajouter  [q] Quitter").
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter)

	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(taskList, 0, 1, true).
		AddItem(footer, 1, 1, false)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'q':
			app.Stop()
		case 'a':
			showAddForm(app, taskCtrl)
		case 'd':
			index := taskList.GetCurrentItem()
			if index >= 0 {
				taskCtrl.MarkDone(index) // méthode à implémenter dans taskCtrl
				refresh()
			}
		}
		return event
	})

	return layout
}

func showAddForm(app *tview.Application, taskCtrl *controller.TaskController) {
	form := tview.NewForm()

	form.AddInputField("Titre", "", 40, nil, nil)

	form.AddButton("Enregistrer", func() {
		title := form.GetFormItem(0).(*tview.InputField).GetText()
		if title != "" {
			task := model.Task{
				Title: title,
				Done:  false,
			}
			taskCtrl.Add(task)
		}
		app.SetRoot(SetupUI(app, taskCtrl), true)
	})

	form.AddButton("Annuler", func() {
		app.SetRoot(SetupUI(app, taskCtrl), true)
	})

	form.SetBorder(true).
		SetTitle(" Nouvelle Tâche ").
		SetTitleAlign(tview.AlignLeft)

	app.SetRoot(form, true).SetFocus(form)
}
