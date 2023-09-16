package doublylinkedlist

func (dll *DoublyLinkedList[DT]) Insert(index int, item DT) (err error) {
	if index == 0 {
		err = dll.Prepend(item)
		return
	}
	if index == dll.length {
		err = dll.Append(item)
		return
	}
	if index < 0 || index > dll.length {
		err = ErrListIndexOutOfRange
		return
	}

	var tempNode *Node[DT]
	newNode := &Node[DT]{
		Data: item,
	}

	if index < dll.length/2 {

		tempNode = dll.head
		for i := 1; i < index; i++ {
			tempNode = tempNode.Next
		}

		newNode.Prev = tempNode
		newNode.Next = tempNode.Next
		newNode.Next.Prev = newNode
		tempNode.Next = newNode
	} else {
		tempNode = dll.tail
		for i := dll.length - 1; i > index; i-- {
			tempNode = tempNode.Prev
		}
		newNode.Next = tempNode
		newNode.Prev = tempNode.Prev
		newNode.Prev.Next = newNode
		tempNode.Prev = newNode
	}

	dll.length++

	return

}
