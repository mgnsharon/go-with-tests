package iteration

import "strings"

// Repeat takes a string and repeats it n times.
func Repeat(char string, times int) (repeated string) {
	for i := 0; i < times; i++ {
		repeated += char
	}
	return repeated
}

// Repeat using std strings package instead of iteration.
func RepeatWStrings(char string, times int) string {
	return strings.Repeat(char, times)
}
