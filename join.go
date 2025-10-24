package piscine

func Join(strs []string, sep string) string {
	if len(strs) == 0 {
		return ""
	}
	res := ""
	for index, str := range strs {
		if index == len(strs)-1 {
			res += str
		} else {
			res += str + sep
		}
	}
	return res
}
