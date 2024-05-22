package main

import (
	"os"
	"time"

	"github.com/mgnsharon/go-with-tests/mocking"
)

func main() {
	sleeper := &mocking.ConfigurableSleeper{Duration: 1 * time.Second, Sleeper: time.Sleep}
	mocking.Countdown(os.Stdout, sleeper)
}
