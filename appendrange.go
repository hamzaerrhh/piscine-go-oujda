package piscine

func AppendRange(min, max int) []int {
	if min >= max {
		return nil
	}
	var slices []int

	for i := min; i < max; i++ {
		slices = append(slices, i)
	}
	return slices
}
