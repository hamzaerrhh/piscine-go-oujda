package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	names := os.Args[1:]
	for i := len(names) - 1; i >= 0; i-- {

		for _, ch := range names[i] {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')

	}
}
