package linked_list

type Node struct {
	Data int
	Next *Node
}

func (node *Node) Add(data int) {
	newNode := Node{data, nil}
	iter := node
	for iter.Next != nil {
		iter = iter.Next
	}
	iter.Next = &newNode
}

func (node *Node) Remove(i int) {
	if node.Data == i {
		*node = *node.Next
	}
	iter := node
	prev := node
	for iter.Next != nil {
		if iter.Data == i {
			prev.Next = iter.Next
		}
		prev = iter
		iter = iter.Next
	}
}
