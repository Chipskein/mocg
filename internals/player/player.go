package player

import (
	"log"
	"os"
	"time"

	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/vorbis"
)

func Play(filepath string) {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := vorbis.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(streamer)
	select {}
}
func Stop() {

}
