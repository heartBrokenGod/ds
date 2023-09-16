package doublylinkedlist

func (sll *DoublyLinkedList[DT]) Lookup(index int) (item DT, err error) {
	// check if list is empty
	if sll.length == 0 {
		err = ErrListIsEmpty
		return
	}
	// check if index greater than equal to length
	if index < 0 || index >= sll.length {
		err = ErrListIndexOutOfRange
		return
	}

	// traverse and return the value
	var tempNode *Node[DT]
	if index < sll.length/2 {
		tempNode = sll.head
		for i := 1; i <= index; i++ {
			tempNode = tempNode.Next
		}
	} else {
		tempNode = sll.tail
		for i := sll.length - 1; i > index; i-- {
			tempNode = tempNode.Prev
		}
	}

	item = tempNode.Data
	return

}
