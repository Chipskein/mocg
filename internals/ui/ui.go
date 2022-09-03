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
	grid             *tui.Grid
	ticker           *<-chan time.Time
	tickerProgresBar *<-chan time.Time
	uiEvents         <-chan tui.Event
	progressBar      *widgets.Gauge
	volumeBar        *widgets.Gauge
	filelist         *widgets.List
	p                *widgets.Paragraph
	repo             *repositories.LocalRepository
	player           *player.PlayerController
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
	volumeBar := widgets.NewGauge()
	volumeBar.TitleStyle.Fg = tui.ColorWhite
	volumeBar.Percent = 50
	volumeBar.Label = "Volume"
	volumeBar.BarColor = tui.ColorWhite
	volumeBar.LabelStyle = tui.NewStyle(tui.ColorWhite)
	t.volumeBar = volumeBar
}
func (t *TUI) RenderProgressBar() {

	processBar := widgets.NewGauge()
	processBar.Title = ""
	processBar.TitleStyle.Fg = tui.ColorWhite
	processBar.Percent = 0
	processBar.Label = ""
	processBar.BarColor = tui.ColorWhite
	processBar.LabelStyle = tui.NewStyle(tui.ColorWhite)
	t.progressBar = processBar
}
func (t *TUI) RenderSongInfo() {
	p := widgets.NewParagraph()
	p.Text = fmt.Sprintf("Time:%s  Duration:%s", "0s", "0s")
	t.p = p
}
func (t *TUI) SetupGrid() {
	grid := tui.NewGrid()
	termWidth, termHeight := tui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	grid.Set(
		tui.NewRow(1.6/2,
			tui.NewCol(2/2,
				t.filelist),
		),
		tui.NewRow(0.2/2,
			tui.NewCol(1.5/2, t.p),
			tui.NewCol(0.5/2, t.volumeBar),
		),
		tui.NewRow(0.18/2, t.progressBar),
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
			case "<Down>", "j":
				t.filelist.ScrollDown()
			case "<Up>", "k":
				t.filelist.ScrollUp()
			case "<End>":
				t.filelist.ScrollBottom()
			case "<Home>":
				t.filelist.ScrollTop()
			case "<PageDown>":
				t.filelist.ScrollHalfPageDown()
			case "<PageUp>":
				t.filelist.ScrollHalfPageUp()
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
				wg.Add(1)
				wg.Done()
				t.volumeBar.Percent = int(t.player.Volume.Volume * 100)
				t.RenderUI()
			case ".":
				go t.player.VolumeUp()
				wg.Add(1)
				wg.Done()
				t.volumeBar.Percent = int(t.player.Volume.Volume * 100)
				t.RenderUI()
			case "m":
				go t.player.Mute()
				wg.Add(1)
				wg.Done()
				if !t.player.Volume.Silent {
					t.volumeBar.Percent = 0
					t.volumeBar.Label = "MUTED"
				} else {
					t.volumeBar.Percent = int(t.player.Volume.Volume * 100)
					t.volumeBar.Label = "Volume"
				}
				t.RenderUI()
			case "<Resize>":
				payload := e.Payload.(tui.Resize)
				t.grid.SetRect(0, 0, payload.Width, payload.Height)
				tui.Clear()
				t.RenderUI()
			}

		case <-*t.ticker:
			t.RenderUI()
		case <-*t.tickerProgresBar:
			if t.progressBar.Percent == 100 {
				t.progressBar.Percent--
			}
			if t.player != nil && !t.player.Ctrl.Paused && t.progressBar.Percent < 100 {
				t.progressBar.Percent++
			}
			t.RenderUI()
		}
		if t.player != nil && !t.player.Ctrl.Paused {
			t.p.Text = fmt.Sprintf("Time:%s  Duration:%s", t.player.Samplerate.D(t.player.Streamer.Position()).Round(time.Second), t.player.Samplerate.D(t.player.Streamer.Len()).Round(time.Second))
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

	t.progressBar.Percent = 0
	t.progressBar.Title = "Playing >"
	t.progressBar.Label = filename
	t.p.Text = fmt.Sprintf("Time:%s  Duration:%s", t.player.Samplerate.D(t.player.Streamer.Position()).Round(time.Second), t.player.Samplerate.D(t.player.Streamer.Len()).Round(time.Second))
	t.volumeBar.Percent = int(t.player.Volume.Volume * 100)
}
func (t *TUI) RenderUI() {
	tui.Render(t.grid)
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
	time.Sleep(time.Millisecond * 100)
	t.SetupGrid()
	t.uiEvents = tui.PollEvents()
	t.ticker = &time.NewTicker(time.Microsecond).C
	t.tickerProgresBar = &time.NewTicker(time.Second).C
	t.HandleTUIEvents()

}
