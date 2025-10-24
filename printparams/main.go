package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	allnames := os.Args[1:]
	for _, name := range allnames {

		for _, ch := range name {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')

	}
}
