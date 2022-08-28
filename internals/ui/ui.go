package ui

import (
	"chipskein/mocg/internals/repositories"
	"fmt"
	"log"
	"sync"
	"time"

	tui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var wg sync.WaitGroup
var TEST_DIR = "" //"../../audios"

func getFileList() []string {
	var key_slice []string
	files, _ := repositories.GetAllFilesFromLocalDirectory(TEST_DIR)
	for key := range files {
		key_slice = append(key_slice, key)
	}
	return key_slice
}

type TUI struct {
	err               error
	grid              *tui.Grid
	ticker            *<-chan time.Time
	tickerProgressBar *<-chan time.Time
	uiEvents          <-chan tui.Event
	progressBar       *widgets.Gauge
	filelist          *widgets.List
	p                 *widgets.Paragraph
	spark             *widgets.SparklineGroup
}

func StartUI() {
	if err := tui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer tui.Close()

	var t = &TUI{}
	go t.RenderFileList()
	go t.RenderVolumeMixer()
	go t.RenderProgressBar()
	go t.RenderSongInfo()
	wg.Add(4)
	wg.Done()
	time.Sleep(time.Millisecond * 5)
	t.SetupGrid()
	t.uiEvents = tui.PollEvents()
	t.ticker = &time.NewTicker(time.Microsecond).C
	t.tickerProgressBar = &time.NewTicker(time.Second / 5).C

	t.HandleTUIEvents()

}

func (t *TUI) RenderFileList() {
	filelist := widgets.NewList()
	filelist.Rows = getFileList()
	filelist.Title = TEST_DIR
	filelist.TitleStyle.Fg = tui.ColorWhite
	filelist.SelectedRowStyle.Fg = tui.ColorBlack
	filelist.SelectedRowStyle.Bg = tui.ColorWhite
	filelist.TextStyle.Fg = tui.ColorMagenta
	t.filelist = filelist
}
func (t *TUI) RenderVolumeMixer() {
	data := []float64{5, 5, 5, 5, 5, 5}
	sl0 := widgets.NewSparkline()
	sl0.Data = data
	sl0.LineColor = tui.ColorMagenta
	sl0.MaxVal = 100
	sl0.Title = "Volume"
	slg0 := widgets.NewSparklineGroup(sl0)
	t.spark = slg0
}
func (t *TUI) RenderProgressBar() {

	processBar := widgets.NewGauge()
	processBar.Title = "Status"
	processBar.TitleStyle.Fg = tui.ColorWhite
	processBar.Percent = 0
	processBar.Label = "Music Name"
	processBar.BarColor = tui.ColorWhite
	processBar.LabelStyle = tui.NewStyle(tui.ColorCyan)
	t.progressBar = processBar
}
func (t *TUI) RenderSongInfo() {
	p := widgets.NewParagraph()
	p.Text = fmt.Sprintf("Filename:%s\n Status:%s\n Time:%s\n Duration:%s\n Loop:%s\n Press H for help\n", "filename.ext", "Playing", "102s", "1m30s", "false")
	t.p = p
}
func (t *TUI) SetupGrid() {
	grid := tui.NewGrid()
	termWidth, termHeight := tui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	grid.Set(
		tui.NewRow(1.8/2,
			tui.NewCol(1.5/2,
				t.filelist),
			tui.NewCol(0.5/2,
				tui.NewRow(0.5/2, t.spark),
				tui.NewRow(0.5/2, t.p),
				tui.NewRow(1.0/2, t.p),
			),
		),
		tui.NewRow(0.2/2,
			t.progressBar),
	)
	t.grid = grid
	t.RenderUI()
}
func (t *TUI) HandleTUIEvents() {
	for {
		select {
		case e := <-t.uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			case "<Enter>":
				//play music file
			case "<Down>":
				t.filelist.ScrollDown()
			case "<Up>":
				t.filelist.ScrollUp()
			case "<Space>":
				//pause
			case ",":
				//volumedown
			case ".":
				//volumeup
			case "h":
				//show help
			case "<Resize>":
				payload := e.Payload.(tui.Resize)
				t.grid.SetRect(0, 0, payload.Width, payload.Height)
				tui.Clear()
				t.RenderUI()
			}
		case <-*t.tickerProgressBar:
			if t.progressBar.Percent == 100 {
				t.progressBar.Percent = 0
			}
			if t.progressBar.Percent < 100 {
				t.progressBar.Percent += 1
			}
		case <-*t.ticker:
			t.RenderUI()
		}
	}
}

func (t *TUI) RenderUI() {
	tui.Render(t.grid)
}
