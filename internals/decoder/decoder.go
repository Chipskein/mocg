package decoder

import (
	"os"

	"github.com/faiface/beep"
	"github.com/faiface/beep/flac"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/vorbis"
	"github.com/faiface/beep/wav"
)

func DecodeOgg(f *os.File) (beep.StreamSeekCloser, beep.Format, error) {
	streamer, format, err := vorbis.Decode(f)
	if err != nil {
		return nil, beep.Format{}, err
	}

	return streamer, format, nil
}
func DecodeFlac(f *os.File) (beep.StreamSeekCloser, beep.Format, error) {
	streamer, format, err := flac.Decode(f)
	if err != nil {
		return nil, beep.Format{}, err
	}
	return streamer, format, nil
}
func DecodeWav(f *os.File) (beep.StreamSeekCloser, beep.Format, error) {
	streamer, format, err := wav.Decode(f)
	if err != nil {
		return nil, beep.Format{}, err
	}
	return streamer, format, nil
}
func DecodeMp3(f *os.File) (beep.StreamSeekCloser, beep.Format, error) {
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		return nil, beep.Format{}, err
	}
	return streamer, format, nil
}
