package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Initialize a slice to store vowels
	newVowel := []rune{}
	args := os.Args[1:] // Skipping the program name

	// If there are no arguments, print a newline and return
	if len(args) == 0 {
		z01.PrintRune('\n')
		return
	}

	// Extract all vowels from the arguments
	for _, argument := range args {
		for _, char := range argument {
			if IsVowel(char) {
				newVowel = append(newVowel, char)
			}
		}
	}

	// Reverse the slice of vowels
	for i, j := 0, len(newVowel)-1; i < j; i, j = i+1, j-1 {
		newVowel[i], newVowel[j] = newVowel[j], newVowel[i]
	}

	// If no vowels were found, just print the arguments as they are
	if len(newVowel) == 0 {
		for i, argument := range args {
			for _, char := range argument {
				z01.PrintRune(char)
			}
			// Print a space between words (except after the last argument)
			if i < len(args)-1 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
		return
	}

	// Replace vowels in the arguments with the reversed vowels from the `newVowel` slice
	vowelIndex := 0
	for i, argument := range args {
		for _, char := range argument {
			if IsVowel(char) {
				// Print the next reversed vowel and move to the next vowel
				z01.PrintRune(newVowel[vowelIndex])
				vowelIndex++
			} else {
				// Print the non-vowel character
				z01.PrintRune(char)
			}
		}
		// Print a space between words (except after the last argument)
		if i < len(args)-1 {
			z01.PrintRune(' ')
		}
	}

	// End the output with a newline
	z01.PrintRune('\n')
}

// Function to check if a character is a vowel (case insensitive)
func IsVowel(char rune) bool {
	vowels := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	for _, v := range vowels {
		if char == v {
			return true
		}
	}
	return false
}
