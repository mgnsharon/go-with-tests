package mocking

import (
	"fmt"
	"io"
	"time"
)

const (
	finalWord = "Go!"
	start     = 3
)

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	Duration time.Duration
	Sleeper  func(time.Duration)
}

func (s *ConfigurableSleeper) Sleep() {
	s.Sleeper(s.Duration)
}

func Countdown(w io.Writer, s Sleeper) {
	for i := start; i > 0; i-- {
		fmt.Fprintln(w, i)
		s.Sleep()
	}

	fmt.Fprint(w, finalWord)
}
