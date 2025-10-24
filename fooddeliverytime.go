package piscine

type food struct {
	Burger, Chips, Nuggets int
}

func FoodDeliveryTime(order string) int {
	food := &food{
		Burger:  15,
		Chips:   10,
		Nuggets: 12,
	}

	if order == "burger" {
		return food.Burger
	}
	if order == "chips" {
		return food.Chips
	}
	if order == "nuggets" {
		return food.Nuggets
	}

	return 404
}
