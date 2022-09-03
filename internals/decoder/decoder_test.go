package decoder

import (
	"sync"
	"testing"
	"time"

	"github.com/chipskein/mocg/internals/player"
	"github.com/chipskein/mocg/internals/repositories"
)

var wg sync.WaitGroup

func TestDecodeOgg(t *testing.T) {
	t.Logf("Play ogg file for 10 seconds")
	var filetest = "../testAudios/ogg/Inamorata-Maruex.ogg"
	f := repositories.ReadFile(filetest)
	streamer, format, err := DecodeOgg(f)
	if err != nil {
		t.Failed()
	}
	defer streamer.Close()
	var p = player.InitPlayer(format.SampleRate, streamer, f)
	go p.Play()
	wg.Add(1)
	time.Sleep(time.Second * 10)
	wg.Done()

	go p.Stop()
	wg.Add(1)
	wg.Done()

	wg.Wait()
}
func TestDecodeMp3(t *testing.T) {
	t.Logf("Play mp3 file for 10 seconds")
	var filetest = "../testAudios/mp3/Inamorata-Maruex.mp3"
	f := repositories.ReadFile(filetest)
	streamer, format, err := DecodeMp3(f)
	if err != nil {
		t.Failed()
	}
	defer streamer.Close()
	var p = player.InitPlayer(format.SampleRate, streamer, f)
	go p.Play()
	wg.Add(1)
	time.Sleep(time.Second * 10)
	wg.Done()

	go p.Stop()
	wg.Add(1)
	wg.Done()

	wg.Wait()
}
func TestDecodeWav(t *testing.T) {
	t.Logf("Play wav file for 10 seconds")
	var filetest = "../testAudios/wav/Inamorata-Maruex.wav"
	f := repositories.ReadFile(filetest)
	streamer, format, err := DecodeWav(f)
	if err != nil {
		t.Failed()
	}
	defer streamer.Close()
	var p = player.InitPlayer(format.SampleRate, streamer, f)
	go p.Play()
	wg.Add(1)
	time.Sleep(time.Second * 10)
	wg.Done()

	go p.Stop()
	wg.Add(1)
	wg.Done()

	wg.Wait()
}
func TestDecodeFlac(t *testing.T) {
	t.Logf("Play flac file for 10 seconds")
	var filetest = "../testAudios/flac/Inamorata-Maruex.flac"
	f := repositories.ReadFile(filetest)
	streamer, format, err := DecodeFlac(f)
	if err != nil {
		t.Failed()
	}
	defer streamer.Close()
	var p = player.InitPlayer(format.SampleRate, streamer, f)
	go p.Play()
	wg.Add(1)
	time.Sleep(time.Second * 10)
	wg.Done()

	go p.Stop()
	wg.Add(1)
	wg.Done()

	wg.Wait()
}
