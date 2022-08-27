package main

import (
	"chipskein/mocg/internals/player"
	"log"
	"os"

	tb "github.com/nsf/termbox-go"
)

func main() {
	err := tb.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer tb.Close()
	var file = "/home/chipskein/Music/Prophecies.ogg"
	for {
		event := tb.PollEvent()
		switch {
		case event.Ch == ',':
			//player.VolumeDown()
		case event.Ch == '.':
			//player.VolumeUp()
		case event.Key == tb.KeyEsc:
			os.Exit(0)
		case event.Key == tb.KeySpace:
			//player.PauseOrResume()
		case event.Key == tb.KeyEnter:
			go player.Play(file)
		case event.Key == tb.KeyArrowUp:
			file = "/home/chipskein/Music/The Perfect Girl.ogg"
		}
	}
}
