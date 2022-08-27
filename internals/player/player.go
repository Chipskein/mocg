package player

import (
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/vorbis"
)

type PlayerController struct {
	ctrl   *beep.Ctrl
	volume *effects.Volume
}

var pctrl PlayerController

func Play(filepath string) {

	if (pctrl != PlayerController{}) {
		speaker.Clear()
		pctrl = PlayerController{}
		return
	}

	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal("could not read file", err)
	}

	streamer, format, err := vorbis.Decode(f)
	if err != nil {
		log.Fatal("Could not decode file", err)
	}

	defer streamer.Close()

	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		log.Fatal("PORRA DE INIT")
	}

	ctrl := &beep.Ctrl{Streamer: beep.Loop(-1, streamer), Paused: false}
	volume := &effects.Volume{
		Streamer: ctrl,
		Base:     2,
		Volume:   0,
		Silent:   false,
	}

	pctrl = PlayerController{ctrl: ctrl, volume: volume}
	speaker.Play(volume)
	for {

	}

}
func PauseOrResume() {
	speaker.Lock()
	pctrl.ctrl.Paused = !pctrl.ctrl.Paused
	speaker.Unlock()
}
func VolumeDown() {
	speaker.Lock()
	pctrl.volume.Volume -= 0.1
	speaker.Unlock()
}
func VolumeUp() {
	speaker.Lock()
	pctrl.volume.Volume += 0.1
	speaker.Unlock()
}
func Stop() {

	//speaker.Clear()
	//pctrl = PlayerController{}

}
