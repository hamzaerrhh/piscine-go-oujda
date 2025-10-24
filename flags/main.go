package main

import (
	"fmt"
	"os"
)

func printHelp() {
	// Printing the help instructions
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("\t This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Println("\t This flag will behave like a boolean, if it is called it will order the argument.")
}

func main() {
	stringToAdd := ""
	stringInsert := ""

	// insert flags
	insertFlag := false
	// orders flags
	orderFlag := false

	// Check for --help or -h
	if len(os.Args) == 1 || (len(os.Args) > 1 && (os.Args[1] == "--help" || os.Args[1] == "-h")) {
		printHelp()
		return
	}

	args := os.Args[1:]

	if len(args) == 0 || len(args) > 3 {
		// No string argument provided
		fmt.Println("")
		return
	}

	for _, argument := range args {
		// Handle --order (or -o) and insert string into the argument
		if argument == "--order" || argument == "-o" {
			orderFlag = true
		} else if len(argument) > 9 && argument[:9] == "--insert=" {
			insertFlag = true
			stringInsert += argument[9:]

		} else if len(argument) > 3 && argument[:3] == "-i=" {
			insertFlag = true
			stringInsert += argument[3:]

		} else {
			stringToAdd += argument
		}
	}
	result := stringToAdd
	// func handle the logic
	if insertFlag {
		result += stringInsert
	}
	if orderFlag {
		result = Order(result)
	}
	// Output the final result
	fmt.Println(result)
}

func Order(str string) string {
	strings := []rune(str)
	// bubble sort
	for i := 0; i < len(strings)-1; i++ {
		for j := i + 1; j < len(strings); j++ {
			if strings[i] > strings[j] {
				strings[i], strings[j] = strings[j], strings[i]
			}
		}
	}
	return string(strings)
}
