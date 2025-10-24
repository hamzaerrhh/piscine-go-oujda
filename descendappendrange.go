package piscine

func DescendAppendRange(max, min int) []int {
	slices := []int{}

	if min >= max {
		return slices
	}
	for max > min {
		slices = append(slices, max)
		max--
	}

	return slices
}
