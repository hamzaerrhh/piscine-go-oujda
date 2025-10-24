package piscine

func ReverseMenuIndex(menu []string) []string {
	newMenu := make([]string, len(menu))

	for i := 0; i < len(menu); i++ {
		newMenu[i] = menu[len(menu)-1-i]
	}
	return newMenu
}
