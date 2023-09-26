package singlylinkedlist

func (sll *SinglyLinkedList[DT]) RemoveAtEnd() (err error) {
	return sll.Remove(sll.length - 1)
}
