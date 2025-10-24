package piscine

func Map(f func(int) bool, a []int) []bool {
	result := make([]bool, 0)

	for _, nbr := range a {
		result = append(result, f(nbr))
	}
	return result
}
