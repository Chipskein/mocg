package player

import (
	"os"
	"sync"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/speaker"
)

type PlayerController struct {
	Ctrl       *beep.Ctrl
	Volume     *effects.Volume
	Samplerate beep.SampleRate
	Streamer   beep.StreamSeekCloser
	Resampler  *beep.Resampler
	Done       *chan bool
	File       *os.File
}

const VOLUME = 0.1
const MAX_VOLUME = 1.0
const MIN_VOLUME = -11.0 //MUTE

/*SEGFAULT when reinit speaker */
var DEFAULT_SAMPLE beep.SampleRate = 48000
var err_speaker = speaker.Init(DEFAULT_SAMPLE, DEFAULT_SAMPLE.N(time.Second/10))

var wg sync.WaitGroup

func (p *PlayerController) Play() {
	speaker.Play(beep.Seq(p.Volume, beep.Callback(func() {
		*p.Done <- true
	})))
	go func() {
	loop:
		for {
			select {
			case <-*p.Done:
				speaker.Clear()
				break loop
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
	p.Ctrl.Paused = !p.Ctrl.Paused
	speaker.Unlock()
}
func (p *PlayerController) VolumeDown() {
	if p.Volume.Volume <= MIN_VOLUME {
		return
	}
	speaker.Lock()
	p.Volume.Volume -= VOLUME
	speaker.Unlock()
}
func (p *PlayerController) VolumeUp() {
	if p.Volume.Volume >= MAX_VOLUME {
		return
	}
	speaker.Lock()
	p.Volume.Volume += VOLUME
	speaker.Unlock()
}
func (p *PlayerController) Stop() {
	*p.Done <- true
}

func InitPlayer(sampleRate beep.SampleRate, streamer beep.StreamSeekCloser, f *os.File) *PlayerController {
	ctrl := &beep.Ctrl{Streamer: beep.Loop(1, streamer)}
	resampler := beep.ResampleRatio(4, 1, ctrl)
	volume := &effects.Volume{Streamer: resampler, Base: 2}
	done := make(chan bool, 1)
	return &PlayerController{Samplerate: sampleRate, Streamer: streamer, Ctrl: ctrl, Resampler: resampler, Volume: volume, Done: &done, File: f}
}
