package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	programeName := []rune(os.Args[0])
	for _, i := range programeName {
		if !(i == '.' || i == '/') {
			z01.PrintRune(i)
		}
	}
	z01.PrintRune('\n')
}
