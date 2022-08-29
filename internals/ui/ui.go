package ui

import (
	"fmt"
	"log"
	"path"
	"sync"
	"time"

	"github.com/chipskein/mocg/internals/decoder"
	"github.com/chipskein/mocg/internals/player"
	"github.com/chipskein/mocg/internals/repositories"

	tui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var wg sync.WaitGroup

type TUI struct {
	grid        *tui.Grid
	ticker      *<-chan time.Time
	uiEvents    <-chan tui.Event
	progressBar *widgets.Gauge
	filelist    *widgets.List
	p           *widgets.Paragraph
	spark       *widgets.SparklineGroup
	repo        *repositories.LocalRepository
	player      *player.PlayerController
}

func StartUI() {
	if err := tui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer tui.Close()

	var t = &TUI{}
	t.repo = &repositories.LocalRepository{CURRENT_DIRECTORY: "../testAudios", DEFAULT_DIRECTORY: "/home/chipskein/Music"}

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

	t.HandleTUIEvents()

}

func (t *TUI) RenderFileList() {
	filelist := widgets.NewList()
	filelist.Rows = t.repo.ListFiles()
	filelist.Title = t.repo.CURRENT_DIRECTORY
	filelist.TitleStyle.Fg = tui.ColorWhite
	filelist.SelectedRowStyle.Fg = tui.ColorBlack
	filelist.SelectedRowStyle.Bg = tui.ColorWhite
	filelist.TextStyle.Fg = tui.ColorWhite
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
				t.HandleSelectedFile(t.filelist.Rows[t.filelist.SelectedRow])
			case "<Down>":
				t.filelist.ScrollDown()
			case "<Up>":
				t.filelist.ScrollUp()
			case "<Space>":
				go t.player.PauseOrResume()
				if !t.player.Ctrl.Paused {
					t.progressBar.Title = "Paused ||"
				} else {
					t.progressBar.Title = "Playing >"
				}
				t.RenderUI()
			case ",":
				go t.player.VolumeDown()
			case ".":
				go t.player.VolumeUp()
			case "<Resize>":
				payload := e.Payload.(tui.Resize)
				t.grid.SetRect(0, 0, payload.Width, payload.Height)
				tui.Clear()
				t.RenderUI()
			}

		case <-*t.ticker:
			t.RenderUI()
		}
	}
}
func (t *TUI) HandleSelectedFile(filename string) {
	if filename == "../" {
		parentPath := path.Dir(t.repo.CURRENT_DIRECTORY)
		t.repo.CURRENT_DIRECTORY = parentPath
		t.filelist.Rows = t.repo.ListFiles()
		t.filelist.Title = t.repo.CURRENT_DIRECTORY
		return
	}
	var file = t.repo.Files[filename]
	if file.IsADirectory {
		t.repo.CURRENT_DIRECTORY = file.FullPath
		t.filelist.Rows = t.repo.ListFiles()
		t.filelist.Title = t.repo.CURRENT_DIRECTORY
		return
	}
	if t.player != nil {
		t.player.Stop()
	}
	var f = repositories.ReadFile(file.FullPath)
	streamer, format, _ := decoder.Decode(f, file.Extension)

	t.player = player.InitPlayer(format.SampleRate, streamer, f)
	go t.player.Play()
	t.progressBar.Title = "Playing >"
	t.progressBar.Label = filename

}
func (t *TUI) RenderUI() {
	tui.Render(t.grid)
}
