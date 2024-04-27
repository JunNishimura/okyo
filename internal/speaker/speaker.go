package speaker

import (
	"bytes"
	"io"
	"log"

	"time"

	"github.com/JunNishimura/okyo/internal/speaker/data"
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
)

type Speaker interface {
	Play()
}

type mp3Speaker struct {
	loopCount int
	speed     float64
}

func NewMp3Speaker(loopCount int, speed float64) Speaker {
	return &mp3Speaker{
		loopCount: loopCount,
		speed:     speed,
	}
}

func (s *mp3Speaker) Play() {
	// not to use beep.Loop method since it doesn't work well when to repeat the file
	// the problem happens since the file is not closed properly(we use io.NopCloser for the embedded file)
	if s.loopCount < 0 {
		s.playInfinite()
	} else {
		s.playFinite()
	}
}

func (s *mp3Speaker) playInfinite() {
	for {
		s.playSound()
	}
}

func (s *mp3Speaker) playFinite() {
	for i := 0; i < s.loopCount; i++ {
		s.playSound()
	}
}

func (s *mp3Speaker) playSound() {
	f := io.NopCloser(bytes.NewReader(data.MP3))
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Printf("fail to decode file: %v", err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	resampledSpeaker := beep.ResampleRatio(4, s.speed, streamer)

	done := make(chan bool)
	speaker.Play(beep.Seq(resampledSpeaker, beep.Callback(func() {
		done <- true
	})))

	<-done
}
