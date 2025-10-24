package piscine

func runeToString(char rune) int {
	return int(char - '0')
}

func BasicAtoi(s string) int {
	nbr := 0
	for _, char := range s {
		// condo remove all the non num value
		if char >= '0' && char <= '9' {
			nbr = runeToString(char) + 10*nbr
		}
	}
	return nbr
}
