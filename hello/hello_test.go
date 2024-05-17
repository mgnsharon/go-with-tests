package hello

import "testing"

func TestHello(t *testing.T) {
	tcs := []struct {
		name     string
		input    string
		expected string
	}{
		{"World", "world", "Hello, world"},
		{"Alice", "Alice", "Hello, Alice"},
		{"Bob", "Bob", "Hello, Bob"},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			actual := Hello(tc.input)
			if actual != tc.expected {
				t.Errorf("Hello(%s) = %s; want %s", tc.input, actual, tc.expected)
			}
		})
	}
}
