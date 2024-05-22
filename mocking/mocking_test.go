package mocking

import (
	"bytes"
	"testing"
	"time"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep(d time.Duration) {
	s.Calls++
}

func TestCountdown(t *testing.T) {

	t.Run("Countdown", func(t *testing.T) {
		b := &bytes.Buffer{}
		spy := &SpySleeper{}

		Countdown(b, spy)

		got := b.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		if spy.Calls != 3 {
			t.Errorf("not enough calls to sleeper, want 3 got %d", spy.Calls)
		}
	})
}
