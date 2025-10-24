package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	slices := make([]int, max-min)

	for i := min; i < max; i++ {
		slices[i-min] = i
	}
	return slices
}
