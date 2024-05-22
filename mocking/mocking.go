package mocking

import (
	"fmt"
	"io"
	"time"
)

const finalWord = "Go!"
const start = 3

type Sleeper interface {
	Sleep(time.Duration)
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep(dur time.Duration) {
	time.Sleep(dur)
}

func Countdown(w io.Writer, s Sleeper) {
	for i := start; i > 0; i-- {
		fmt.Fprintln(w, i)
		s.Sleep(1 * time.Second)
	}
	fmt.Fprint(w, finalWord)

}
