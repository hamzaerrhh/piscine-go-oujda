package piscine

func ListSize(l *List) int {
	count := 0
	node := l.Head
	if l.Head == nil {
		return 0
	}
	for node != nil {
		count++
		node = node.Next
	}
	return count
}
