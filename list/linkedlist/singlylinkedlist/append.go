package singlylinkedlist

func (sll *SinglyLinkedList[DT]) Append(item DT) error {

	// create new node
	newNode := &Node[DT]{
		Data: item,
		Next: nil,
	}

	// if singly link list is empty
	if sll.length == 0 {
		sll.head = newNode
		sll.tail = newNode
		sll.length++
		return nil
	}

	// if not empty
	sll.tail.Next = newNode
	sll.tail = newNode
	sll.length++

	return nil
}
