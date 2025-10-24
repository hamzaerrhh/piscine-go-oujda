package piscine

func BasicAtoi2(s string) int {
	nbr := 0
	for _, char := range s {
		// u should add a funcion return 0 if a string

		// condo remove all the non num value
		if char >= '0' && char <= '9' {
			nbr = int(char-'0') + 10*nbr
		} else {
			return 0
		}
	}
	return nbr
}
