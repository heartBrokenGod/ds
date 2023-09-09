package singlylinkedlist

func (sll *SinglyLinkedList[DT]) Lookup(index int) (item DT, err error) {
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
	tempNode := sll.head
	for i := 1; i <= index; i++ {
		tempNode = tempNode.Next
	}

	item = tempNode.Data
	return

}
