package piscine

func StringToIntSlice(str string) []int {
	var slice []int

	if str == "" {
		return slice
	}
	for _, char := range str {
		slice = append(slice, int(char))
	}
	return slice
}
