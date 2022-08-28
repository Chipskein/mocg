package decoder

import (
	"chipskein/mocg/internals/player"
	"chipskein/mocg/internals/repositories"
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup

func TestDecodeOgg(t *testing.T) {
	var filetest = "../testAudios/ogg/music3.ogg"
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
	var filetest = "../testAudios/mp3/music3.mp3"
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
	var filetest = "../testAudios/wav/music3.wav"
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
	var filetest = "../testAudios/flac/music3.flac"
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
