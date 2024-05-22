package main

import (
	"os"

	"github.com/mgnsharon/go-with-tests/mocking"
)

func main() {
	sleeper := &mocking.DefaultSleeper{}
	mocking.Countdown(os.Stdout, sleeper)
}
