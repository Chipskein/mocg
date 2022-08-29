package player

import (
	"testing"
	"time"

	r "github.com/chipskein/mocg/internals/repositories"

	"github.com/faiface/beep/vorbis"
)

const MUSIC_TIMEOUT = time.Second * 10

func TestPlay(T *testing.T) {
	var filename = "../testAudios/ogg/music1.ogg"
	T.Logf("Play a song file for %d seconds them stop", MUSIC_TIMEOUT)
	var f = r.ReadFile(filename)
	streamer, format, err := vorbis.Decode(f)
	if err != nil {
		T.Fatalf("Could not Decode %s ERROR %s", filename, err)
	}
	defer streamer.Close()

	var p = InitPlayer(format.SampleRate, streamer, f)
	go p.Play()
	wg.Add(1)
	time.Sleep(MUSIC_TIMEOUT)
	wg.Done()

	go p.Stop()
	wg.Add(1)
	wg.Done()

	wg.Wait()
}

func TestPlay2(T *testing.T) {
	var filename = "../testAudios/ogg/music2.ogg"
	T.Logf("Play a song file for %d seconds them stop", MUSIC_TIMEOUT)

	var f = r.ReadFile(filename)
	streamer, format, err := vorbis.Decode(f)
	if err != nil {
		T.Fatalf("[ERROR] Could not decode %s Error %s", filename, err)
	}
	defer streamer.Close()

	var p = InitPlayer(format.SampleRate, streamer, f)
	go p.Play()
	wg.Add(1)
	time.Sleep(MUSIC_TIMEOUT)
	wg.Done()

	go p.Stop()
	wg.Add(1)
	wg.Done()

	wg.Wait()

}
func TestVolume(T *testing.T) {
	var filename = "../testAudios/ogg/music3.ogg"
	var MUSIC_TIMEOUT = time.Second * 20
	var MAX_VOLUME_UP = 35
	var MAX_VOLUME_DOWN = 50

	T.Logf("Play a song file for %d seconds and change volume up and down them stop", MUSIC_TIMEOUT)

	var f = r.ReadFile(filename)
	streamer, format, err := vorbis.Decode(f)
	if err != nil {
		T.Fatalf("[ERROR] Could not decode %s Error %s", filename, err)
	}
	defer streamer.Close()

	var p = InitPlayer(format.SampleRate, streamer, f)
	go p.Play()
	wg.Add(1)
	wg.Done()
	time.Sleep(time.Second * 2)
	go func() {
		for i := 0; i <= MAX_VOLUME_UP; i++ {
			go p.VolumeUp()
			wg.Add(1)
			wg.Done()
			time.Sleep(time.Millisecond * 50)
		}
	}()

	time.Sleep(time.Second * 4)

	go func() {
		for i := 0; i <= MAX_VOLUME_DOWN; i++ {
			go p.VolumeDown()
			wg.Add(1)
			wg.Done()
			time.Sleep(time.Millisecond * 50)
		}
	}()

	time.Sleep(MUSIC_TIMEOUT)
	go p.Stop()
	wg.Add(1)
	wg.Done()

	wg.Wait()
}
func TestPause(T *testing.T) {
	var filename = "../testAudios/ogg/music1.ogg"
	T.Logf("Play a song file for %d seconds and pause/resume during execution", MUSIC_TIMEOUT)

	var f = r.ReadFile(filename)
	streamer, format, err := vorbis.Decode(f)
	if err != nil {
		T.Fatalf("[ERROR] Could not decode %s Error %s", filename, err)
	}
	defer streamer.Close()

	var p = InitPlayer(format.SampleRate, streamer, f)
	go p.Play()
	wg.Add(1)
	wg.Done()
	T.Logf("Pause")
	time.Sleep(time.Second * 4)
	p.PauseOrResume()

	T.Logf("Resume")
	time.Sleep(time.Second * 4)
	p.PauseOrResume()

	time.Sleep(MUSIC_TIMEOUT)
	go p.Stop()
	wg.Add(1)
	wg.Done()

	wg.Wait()
}
