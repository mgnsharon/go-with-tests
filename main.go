package main

import (
	"fmt"

	"github.com/mgnsharon/go-with-tests/hello"
)

func main() {
	fmt.Println("Learn Go with Tests")

	// Separate the domain code from the side effects
	// The message is the domain, the printing is the side effect
	fmt.Println(hello.Hello("Megan", hello.English))
}
