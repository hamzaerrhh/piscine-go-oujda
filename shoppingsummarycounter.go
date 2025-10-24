package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	grocery := make(map[string]int, 0)

	slices := FixSlices(str, " ")

	// make the map
	for _, value := range slices {
		grocery[value]++
	}

	return grocery
	// add the values
}

func FixSlices(s string, sep string) []string {
	var result []string
	word := ""
	sepLen := len(sep)
	for i := 0; i < len(s); {
		if i+sepLen <= len(s) && s[i:i+sepLen] == sep {
			result = append(result, word)
			word = ""
			i += sepLen
		} else {
			word += string(s[i])
			i++
		}
	}

	result = append(result, word)
	return result
}
