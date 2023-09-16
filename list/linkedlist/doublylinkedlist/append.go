package doublylinkedlist

func (dll *DoublyLinkedList[DT]) Append(item DT) error {

	newNode := &Node[DT]{
		Data: item,
		Prev: nil,
		Next: nil,
	}
	// check if list is empty
	if dll.length == 0 {
		dll.head = newNode
		dll.tail = newNode
		dll.length++
		return nil
	}

	newNode.Prev = dll.tail
	dll.tail.Next = newNode
	dll.tail = newNode
	dll.length++

	return nil
}
