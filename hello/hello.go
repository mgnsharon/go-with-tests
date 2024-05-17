package hello

import "fmt"

type Language string

const (
	englishHelloPrefix = "Hello,"
	spanishHelloPrefix = "Hola,"
	frenchHelloPrefix  = "Bonjour,"

	English Language = "English"
	Spanish Language = "Spanish"
	French  Language = "French"
)

var greetings = map[Language]string{
	English: englishHelloPrefix,
	Spanish: spanishHelloPrefix,
	French:  frenchHelloPrefix,
}

func Hello(n string, l Language) string {
	if n == "" {
		n = "World"
	}
	if l == "" {
		l = English
	}
	return fmt.Sprintf("%s %s", greetings[l], n)
}
