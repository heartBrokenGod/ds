package doublylinkedlist

func (dll *DoublyLinkedList[DT]) Prepend(item DT) error {

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

	newNode.Next = dll.head
	dll.head.Prev = newNode
	dll.head = newNode

	return nil
}
