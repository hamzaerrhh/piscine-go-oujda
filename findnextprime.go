package piscine

func FindNextPrime(nb int) int {
	if nb < 0 {
		return 2
	}
	for i := nb; ; i++ {
		if IsPrime(i) {
			return i
		}
	}
	return 0
}
