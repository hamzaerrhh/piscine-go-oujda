package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	args := os.Args[1:]
	for _, str := range args {
		if str == "01" || str == "galaxy" || str == "galaxy 01" {
			fmt.Println("Alert!!!")
			return
		}
	}
}
