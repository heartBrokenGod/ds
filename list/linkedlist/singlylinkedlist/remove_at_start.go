package singlylinkedlist

func (sll *SinglyLinkedList[DT]) RemoveAtStart() (err error) {
	return sll.Remove(0)
}
