package hello

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	tcs := []struct {
		name     string
		input    string
		lang     Language
		expected string
	}{
		{"Valid Name", "Megan", "", "Hello, Megan"},
		{"Empty String", "", "", "Hello, World"},
		{"In Spanish", "Elodie", Spanish, "Hola, Elodie"},
		{"In French", "Jean", French, "Bonjour, Jean"},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			actual := Hello(tc.input, tc.lang)
			if actual != tc.expected {
				t.Errorf("Hello(%s, %s) = %s; want %s", tc.input, tc.lang, actual, tc.expected)
			}
		})
	}
}

func ExampleHello() {
	greeting := Hello("Jesus", Spanish)
	fmt.Println(greeting)
	// Output: Hola, Jesus
}
