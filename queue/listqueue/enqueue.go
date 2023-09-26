package listqueue

func (listQueue *ListQueue[DT]) Enqueue(item DT) error {
	err := listQueue.singlyLinkedList.Append(item)
	return err
}
