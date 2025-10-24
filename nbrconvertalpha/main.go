package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	isUpper := false

	for _, v := range arguments {
		if v == "--upper" {
			arguments = arguments[1:]
			isUpper = true
		}
	}

	for _, arg := range arguments {
		numv := 0
		for _, v := range arg {
			numv = numv*10 + int(v-'0')
		}

		if numv >= 1 && numv <= 26 {
			if !isUpper {
				z01.PrintRune(rune(numv + 'a' - 1))
			} else {
				z01.PrintRune(rune(numv + 'A' - 1))
			}
		} else {
			z01.PrintRune(' ')
		}
	}
	if len(arguments) > 1 {
		z01.PrintRune('\n')
	}
}
