package decoder

import (
	"os"

	"github.com/faiface/beep"
	"github.com/faiface/beep/vorbis"
)

func DecodeOgg(f *os.File) (beep.StreamSeekCloser, beep.Format, error) {
	streamer, format, err := vorbis.Decode(f)
	if err != nil {
		return nil, beep.Format{}, err
	}
	defer streamer.Close()
	return streamer, format, nil
}
func DecodeFlac() {}
func DecodeWav()  {}
func DecodeMp3()  {}
