package piscine

func Abort(a, b, c, d, e int) int {
	slices := [5]int{a, b, c, d, e}

	for i := 0; i < 4; i++ {
		for j := i + 1; j < 5; j++ {
			if slices[i] > slices[j] {
				slices[i], slices[j] = slices[j], slices[i]
			}
		}
	}

	return slices[2]
}
