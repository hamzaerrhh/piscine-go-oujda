package piscine

func SortIntegerTable(table []int) {
	for i := 0; i < len(table)-1; i++ {
		for j := 0; j < len(table)-1-i; j++ {
			if table[j] > table[j+1] {
				swap(&table[j], &table[j+1])
			}
		}
	}
}

func swap(a *int, b *int) {
	*a, *b = *b, *a
}
