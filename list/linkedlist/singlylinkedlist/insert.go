package singlylinkedlist

func (sll *SinglyLinkedList[DT]) Insert(index int, item DT) (err error) {
	if index == 0 {
		err = sll.Prepend(item)
		return
	}
	if index == sll.length {
		err = sll.Append(item)
		return
	}
	if index < 0 || index > sll.length {
		err = ErrListIndexOutOfRange
		return
	}

	tempNode := sll.head
	for i := 1; i < index; i++ {
		tempNode = tempNode.Next
	}

	newNode := &Node[DT]{
		Data: item,
		Next: tempNode.Next,
	}

	tempNode.Next = newNode

	sll.length++

	return

}
