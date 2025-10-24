package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	assending := false

	if len(a) > 2 {
		if f(a[0], a[1]) > 0 {
			assending = true
		} else if f(a[0], a[1]) < 0 {
			assending = false
		}
		// else {
		// 	return false
		// }
	}

	for i := 1; i < len(a)-1; i++ {
		// case of assending number
		if !assending && f(a[i], a[i+1]) > 0 {
			return false
		}
		if assending && f(a[i], a[i+1]) < 0 {
			return false
		}

		// if f(a[i], a[i+1]) == 0 {
		// 	return false
		// }
	}

	return true
}
