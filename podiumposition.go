package piscine

func PodiumPosition(podium [][]string) [][]string {
	result := podium
	for i := 0; i < len(result)/2; i++ {
		j := len(result) - 1 - i
		result[i], result[j] = result[j], result[i]
	}

	return result
}
