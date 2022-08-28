package ui

import (
	"github.com/rivo/tview"
)

type UI struct {
	app     *tview.Application
	list    *tview.List
	mainBox *tview.Box
}

var tui = UI{}

func renderList() *tview.List {
	return renderList()
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
