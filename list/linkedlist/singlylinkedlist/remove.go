package singlylinkedlist

func (sll *SinglyLinkedList[DT]) Remove(index int) (err error) {

	// check if list is empty
	if sll.IsEmpty() {
		return ErrListIsEmpty
	}
	// check if index is valid
	if index < 0 || index > sll.length-1 {
		err = ErrListIndexOutOfRange
		return
	}
	// remove at start
	if index == 0 {
		sll.head = sll.head.Next
		sll.length--
		// if only one element
		if sll.length == 0 {
			sll.tail = nil
		}
		return
	}

	// remove at any index
	temp := sll.head

	for i := 0; i < index-1; i++ {
		temp = temp.Next
	}
	temp.Next = temp.Next.Next

	sll.length--
	if temp.Next == nil {
		sll.tail = temp
	}
	return

}
