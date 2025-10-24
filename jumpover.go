package piscine

func JumpOver(str string) string {
	if len(str) < 3 {
		return "\n"
	}
	res := ""
	for index, char := range str {
		if (index+1)%3 == 0 {
			res += string(char)
		}
	}

	return res + "\n"
}
