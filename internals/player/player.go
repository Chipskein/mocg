package player

import (
	"fmt"
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

var (
	pctrl PlayerController
	done  = make(chan bool, 100)
)

func Play(filepath string) {
	fmt.Println(filepath)
	if (pctrl != PlayerController{}) {
		done <- true
	}

	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal("[ERROR] Could not read file", err)
	}

	streamer, format, err := vorbis.Decode(f)
	if err != nil {
		log.Fatal("[ERROR] Could not decode file", err)
	}
	defer streamer.Close()

	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	if err != nil {
		log.Fatalln("[ERROR] Acessing speaker")
	}

	ctrl := &beep.Ctrl{Streamer: beep.Loop(1, streamer), Paused: false}
	volume := &effects.Volume{
		Streamer: ctrl,
		Base:     2,
		Volume:   0,
		Silent:   false,
	}

	pctrl = PlayerController{ctrl: ctrl, volume: volume}
	speaker.Play(beep.Seq(pctrl.volume))

	for {
		select {
		case <-done:
			streamer.Close()
			speaker.Clear()
			log.Println("SIGNAL ON DONE CHANNEL")
		}
	}

}
func PauseOrResume() {
	if (pctrl != PlayerController{}) {
		speaker.Lock()
		pctrl.ctrl.Paused = !pctrl.ctrl.Paused
		speaker.Unlock()
	}
}
func VolumeDown() {
	if (pctrl != PlayerController{}) {
		speaker.Lock()
		pctrl.volume.Volume -= 0.1
		speaker.Unlock()
	}
}
func VolumeUp() {
	if (pctrl != PlayerController{}) {
		speaker.Lock()
		pctrl.volume.Volume += 0.1
		speaker.Unlock()
	}
}
