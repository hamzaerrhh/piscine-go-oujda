package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	// declare the new nodeL
	newNode := &NodeL{Data: data}

	// if my new list has no head and tail
	if l.Head == nil {

		l.Head = newNode
		l.Tail = newNode
		return
	} else {
		l.Tail.Next = newNode
		l.Tail = newNode
	}
}
