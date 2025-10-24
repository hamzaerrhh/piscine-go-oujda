package piscine

func Compact(ptr *[]string) int {
	count := 0
	newElem := make([]string, 0)
	for _, elem := range *ptr {
		if elem != "" {
			count++
			newElem = append(newElem, elem)
		}
	}
	*ptr = newElem
	return count
}
