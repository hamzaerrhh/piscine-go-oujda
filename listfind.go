package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	if l == nil || l.Head == nil {
		return nil
	}
	current := l.Head

	for l.Head != nil {
		if comp(l.Head.Data, ref) {
			return &l.Head.Data
		}
		l.Head = l.Head.Next
	}
	l.Head = current
	return nil
}
