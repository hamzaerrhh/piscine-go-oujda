package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	mergedList := merge(n1, n2)
	return ListSort(mergedList)
}

func merge(n1 *NodeI, n2 *NodeI) *NodeI {
	if n1 == nil || n2 == nil {
		return nil
	}

	if n1 == nil {
		return n2
	}

	if n2 == nil {
		return n2
	}
	// go to next is nil and it equal to n2
	current := n1
	for {
		if current.Next == nil {
			current.Next = n2
			break
		}
		current = current.Next

	}
	return n1
}
