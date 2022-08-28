package repositories

import (
	"chipskein/mocg/internals/player"
	"sync"
	"testing"
	"time"

	"github.com/faiface/beep/vorbis"
)

const TEST_DIRECTORY = "../../audios"
const MUSIC_TIMEOUT = time.Second * 5

var wg sync.WaitGroup

func TestReadLocalDirectory(t *testing.T) {
	t.Logf("Try to read all files from directory %s", TEST_DIRECTORY)
	files, _ := GetAllFilesFromLocalDirectory(TEST_DIRECTORY)
	for key := range files {
		t.Log(key)
	}
}
func TestReadLocalDirectory2(t *testing.T) {
	var TEST_DIRECTORY = "/fasfasf/as"
	t.Logf("Try to read all files from directory %s , Should use DEFAULT DIRECTORY", TEST_DIRECTORY)
	files, _ := GetAllFilesFromLocalDirectory(TEST_DIRECTORY)
	for key := range files {
		t.Log(key)
	}
}

func TestReadLocalDirectory3(t *testing.T) {
	var TEST_DIRECTORY = "/fasfasf/as"
	DEFAULT_DIRECTORY = TEST_DIRECTORY
	t.Logf("Try to read all files from directory %s with wrong DEFAULT_DIRECTORY directory\n SHOULD FAIL \n", TEST_DIRECTORY)
	_, err := GetAllFilesFromLocalDirectory(TEST_DIRECTORY)
	if err != nil {
		t.Fatal(err)
	}

}
func TestIntegrationWithPlayer(t *testing.T) {
	t.Logf("Foreach file in %s play a audio for 5 seconds", TEST_DIRECTORY)
	files, _ := GetAllFilesFromLocalDirectory(TEST_DIRECTORY)
	for _, file := range files {
		var path = file.FullPath
		t.Logf("Play a song file for %d seconds them stop", MUSIC_TIMEOUT)
		var f = player.ReadFile(path)
		streamer, format, err := vorbis.Decode(f)
		if err != nil {
			t.Fatalf("Could not Decode %s ERROR %s", path, err)
		}
		defer streamer.Close()

		var p = player.InitPlayer(format.SampleRate, streamer, f)
		go p.Play()
		wg.Add(1)
		time.Sleep(MUSIC_TIMEOUT)
		wg.Done()

		go p.Stop()
		wg.Add(1)
		wg.Done()

		wg.Wait()
	}
}
