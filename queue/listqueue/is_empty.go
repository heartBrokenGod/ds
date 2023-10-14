package listqueue

func (listQueue *ListQueue[DT]) IsEmpty() bool {
	return listQueue.singlyLinkedList.IsEmpty()
}
