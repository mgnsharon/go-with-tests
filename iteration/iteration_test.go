package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	tcs := []struct {
		name     string
		char     string
		times    int
		expected string
	}{
		{"Repeat 5 times", "a", 5, "aaaaa"},
		{"Repeat 3 times", "b", 3, "bbb"},
		{"Repeat 0 times", "c", 0, ""},
		{"Repeat 1 times", "d", 1, "d"},
		{"Repeat 2 times", "e", 2, "ee"},
		{"Send more than 1 char", "fz06", 2, "fz06fz06"},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			actual := Repeat(tc.char, tc.times)
			if actual != tc.expected {
				t.Errorf("Repeat(%s, %d) = %s; want %s", tc.char, tc.times, actual, tc.expected)
			}
		})
	}
}

func ExampleRepeat() {
	repeated := Repeat("z", 3)
	fmt.Println(repeated)
	// Output: zzz
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("z", 3)
	}
}

func TestRepeatWithStrings(t *testing.T) {
	tcs := []struct {
		name     string
		char     string
		times    int
		expected string
	}{
		{"Repeat 5 times", "a", 5, "aaaaa"},
		{"Repeat 3 times", "b", 3, "bbb"},
		{"Repeat 0 times", "c", 0, ""},
		{"Repeat 1 times", "d", 1, "d"},
		{"Repeat 2 times", "e", 2, "ee"},
		{"Send more than 1 char", "fz06", 2, "fz06fz06"},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			actual := Repeat(tc.char, tc.times)
			if actual != tc.expected {
				t.Errorf("Repeat(%s, %d) = %s; want %s", tc.char, tc.times, actual, tc.expected)
			}
		})
	}
}

func BenchmarkRepeatWithStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("z", 3)
	}
}
