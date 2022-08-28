package player

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/speaker"
)

type PlayerController struct {
	ctrl       *beep.Ctrl
	volume     *effects.Volume
	samplerate beep.SampleRate
	streamer   beep.StreamSeekCloser
	resampler  *beep.Resampler
	done       *chan bool
	file       *os.File
}

const VOLUME = 0.1

/*SEGFAULT when reinit speaker */
var DEFAULT_SAMPLE beep.SampleRate = 48000
var err_speaker = speaker.Init(DEFAULT_SAMPLE, DEFAULT_SAMPLE.N(time.Second/10))

var wg sync.WaitGroup

func (p *PlayerController) Play() {
	speaker.Play(beep.Seq(p.volume, beep.Callback(func() {
		*p.done <- true
	})))
	go func() {
	loop:
		for {
			select {
			case <-*p.done:
				speaker.Clear()
				break loop
			case <-time.After(time.Second):
				fmt.Println(p.file.Name(), p.samplerate.D(p.streamer.Position()).Round(time.Second))
			}

		}
	}()
	wg.Add(1)
	wg.Done()
	wg.Wait()
	return
}
func (p *PlayerController) PauseOrResume() {
	speaker.Lock()
	p.ctrl.Paused = !p.ctrl.Paused
	speaker.Unlock()
}
func (p *PlayerController) VolumeDown() {
	speaker.Lock()
	p.volume.Volume -= VOLUME
	speaker.Unlock()
}
func (p *PlayerController) VolumeUp() {
	speaker.Lock()
	p.volume.Volume += VOLUME
	speaker.Unlock()
}
func (p *PlayerController) Stop() {
	*p.done <- true
}

func InitPlayer(sampleRate beep.SampleRate, streamer beep.StreamSeekCloser, f *os.File) *PlayerController {
	ctrl := &beep.Ctrl{Streamer: beep.Loop(1, streamer)}
	resampler := beep.ResampleRatio(4, 1, ctrl)
	volume := &effects.Volume{Streamer: resampler, Base: 2}
	done := make(chan bool, 1)
	return &PlayerController{samplerate: sampleRate, streamer: streamer, ctrl: ctrl, resampler: resampler, volume: volume, done: &done, file: f}
}
