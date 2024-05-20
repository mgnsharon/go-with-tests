package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {

	t.Run("Greet Buffer", func(t *testing.T) {
		b := bytes.Buffer{}
		Greet(&b, "Megan")

		got := b.String()
		want := "Hello, Megan"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
