package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	params := os.Args[1:]

	// sort params
	for i := 0; i < len(params)-1; i++ {
		for j := i + 1; j < len(params); j++ {
			if params[i] > params[j] {
				params[i], params[j] = params[j], params[i]
			}
		}
	}

	// printparams
	for _, char := range params {
		for _, ch := range char {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')

	}
}
