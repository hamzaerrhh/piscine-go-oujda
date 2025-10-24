package piscine

import (
	"fmt"
)

func DealAPackOfCards(deck []int) {
	count := 0
	index := 0
	for i := 0; i <= 3; i++ {

		count++
		fmt.Printf("Player %d: %d, %d, %d\n", count, deck[index], deck[index+1], deck[index+2])
		index += 3
	}
}
