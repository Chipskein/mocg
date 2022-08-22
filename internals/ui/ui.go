package ui

import (
	"chipskein/mocg/internals/repositories"

	"github.com/rivo/tview"
)

type UI struct {
	app     *tview.Application
	list    *tview.List
	mainBox *tview.Box
}

var tui = UI{}

func renderList() *tview.List {
	files := repositories.GetAllFilesFromLocalDirectory("")
	list := tview.NewList()
	for filename, file := range files {
		var Iteration_file = file
		list.AddItem(filename, file.FullPath, '*', func() {
			repositories.HandleFile(Iteration_file)
		})
	}
	list.AddItem("Quit", "Press to exit", 'q', KillUI)
	return list
}

func StartUI() {

	tui.app = tview.NewApplication()
	tui.list = renderList()
	if err := tui.app.SetRoot(tui.list, true).Run(); err != nil {
		panic(err)
	}
}
func KillUI() {
	tui.app.Stop()
}
