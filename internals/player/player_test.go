package player

import (
	"testing"
	"time"
)

func TestPlay(t *testing.T) {
	go Play("/home/chipskein/Music/The Perfect Girl.ogg")
	time.Sleep(time.Second * 5)
	go Play("/home/chipskein/Music/Стыд - Одинокий гражданин.ogg")
	for {

	}
}
