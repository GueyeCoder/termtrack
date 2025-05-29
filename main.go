package main

import (
	"log"

	"github.com/GueyeCoder/termtrack/controller"
	"github.com/GueyeCoder/termtrack/ui"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	taskCtrl, err := controller.NewTaskController("data/tasks.json")
	if err != nil {
		log.Fatalf("Erreur chargement tâches : %v", err)
	}

	root := ui.SetupUI(app, taskCtrl)

	if err := app.SetRoot(root, true).Run(); err != nil {
		panic(err)
	}
}
