package piscine

func ConcatParams(args []string) string {
	if len(args) < 1 {
		return "\n"
	}
	res := ""
	for index, arg := range args {
		if index == len(args)-1 {
			res += arg
		} else {
			res += arg + "\n"
		}
	}
	return res
}
