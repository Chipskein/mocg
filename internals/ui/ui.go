package ui

import (
	"chipskein/mocg/internals/repositories"
	"fmt"
	"log"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var TEST_DIR = "" //"../../audios"

func getFileList() []string {
	var key_slice []string
	files, _ := repositories.GetAllFilesFromLocalDirectory(TEST_DIR)
	for key := range files {
		key_slice = append(key_slice, key)
	}
	return key_slice
}

func StartUI() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	filelist := widgets.NewList()
	filelist.Rows = getFileList()
	filelist.Title = TEST_DIR
	filelist.TitleStyle.Fg = ui.ColorWhite
	filelist.SelectedRowStyle.Fg = ui.ColorBlack
	filelist.SelectedRowStyle.Bg = ui.ColorWhite
	filelist.TextStyle.Fg = ui.ColorMagenta

	data := []float64{4, 2, 1, 6, 3, 9, 1, 4, 2, 15, 14, 9, 8, 6, 10, 13, 15, 12, 10, 5, 3, 6, 1, 7, 10, 10, 14, 13, 6}

	sl0 := widgets.NewSparkline()
	sl0.Data = data[3:]
	sl0.LineColor = ui.ColorGreen

	// single
	slg0 := widgets.NewSparklineGroup(sl0)

	p := widgets.NewParagraph()
	p.Text = fmt.Sprintf("Filename:%s\n Status:%s\n Time:%s\n Duration:%s\n Loop:%s\n Press H for help\n", "filename.ext", "Playing", "102s", "1m30s", "false")

	processBar := widgets.NewGauge()
	processBar.Title = "Status"
	processBar.TitleStyle.Fg = ui.ColorWhite
	processBar.Percent = 0
	processBar.Label = "Music Name"
	processBar.BarColor = ui.ColorWhite
	processBar.LabelStyle = ui.NewStyle(ui.ColorCyan)

	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	grid.Set(
		ui.NewRow(1.8/2,
			ui.NewCol(1.5/2, filelist),
			ui.NewCol(0.5/2,
				ui.NewRow(1.0/2, slg0),
				ui.NewRow(1.0/2, p),
			),
		),
		ui.NewRow(0.2/2, processBar),
	)

	ui.Render(grid)

	tickerCount := 1
	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(time.Second / 5).C
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			case "<Enter>":
				//play music file
			case "<Down>":
				filelist.ScrollDown()
			case "<Up>":
				filelist.ScrollUp()
			case "<Space>":
				//pause
			case ",":
				//volumedown
			case ".":
				//volumeup
			case "h":
				//show help
			case "<Resize>":
				payload := e.Payload.(ui.Resize)
				grid.SetRect(0, 0, payload.Width, payload.Height)
				ui.Clear()
				ui.Render(grid)
			}
		case <-ticker:
			if processBar.Percent == 100 {
				processBar.Percent = 0
			}
			if processBar.Percent < 100 {
				processBar.Percent += 1
			}

			ui.Render(grid)
			tickerCount++
		}
	}
}
