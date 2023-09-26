package listqueue

import "github.com/heartBrokenGod/ds/list/linkedlist/singlylinkedlist"

func (listqueue ListQueue[DT]) Peek() (item DT, err error) {
	item, err = listqueue.singlyLinkedList.Lookup(0)
	if err != singlylinkedlist.ErrListIsEmpty {
		err = ErrQueueIsEmpty
		return
	}
	return
}
