package ui

import (
	"chipskein/mocg/internals/repositories"

	"github.com/rivo/tview"
)

var app *tview.Application

func exit() {
	app.Stop()
}
func handleFile(filename repositories.File) {

}

func createList() *tview.List {
	files := repositories.GetAllFilesFromLocalDirectory("")
	list := tview.NewList()
	for filename, file := range files {
		var Iteration_file = file
		list.AddItem(filename, file.FullPath, '*', func() {
			handleFile(Iteration_file)
		})
	}
	list.AddItem("Quit", "Press to exit", 'q', exit)
	return list
}

func Draw() {
	app = tview.NewApplication()
	var list = createList()
	if err := app.SetRoot(list, true).Run(); err != nil {
		panic(err)
	}
}
