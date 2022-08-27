package main

import (
	"chipskein/mocg/internals/player"
	"fmt"

	tb "github.com/nsf/termbox-go"
)

func main() {
	err := tb.Init()
	if err != nil {
		panic(err)
	}
	defer tb.Close()
	for {
		fmt.Println("TESTANDO")
		event := tb.PollEvent()
		switch {
		case event.Ch == ',':
			go player.VolumeDown()
		case event.Ch == '.':
			go player.VolumeUp()
		case event.Key == tb.KeyEsc:
			player.Stop()
		case event.Key == tb.KeySpace:
			go player.PauseOrResume()
		case event.Key == tb.KeyEnter:
			go player.Play("/home/chipskein/Music/Prophecies.ogg")
		}
	}
}
