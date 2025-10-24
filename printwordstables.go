package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	if len(a) < 1 {
		z01.PrintRune('\n')
		return
	}
	for _, text := range a {

		for _, char := range text {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')

	}
}
