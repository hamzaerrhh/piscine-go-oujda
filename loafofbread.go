package piscine

func LoafOfBread(str string) string {
	res := ""
	world := ""
	if str == "" {
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Output\n"
	}
	for _, val := range str {
		if val != ' ' && len(world) != 5 {
			world += string(val)
		} else if len(world) == 5 {
			if res != "" {
				res += " "
			}
			res += world
			world = ""
		}
	}
	if world != "" {
		if res != "" {
			res += " "
		}
		res += world
	}
	return res + "\n"
}
