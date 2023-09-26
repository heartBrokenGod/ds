package listqueue

import (
	"errors"

	"github.com/heartBrokenGod/ds/list/linkedlist/singlylinkedlist"
)

func (listQueue *ListQueue[DT]) Dequeue() (item DT, err error) {
	item, err = listQueue.Peek()
	if err != nil && errors.Is(err, singlylinkedlist.ErrListIsEmpty) {
		err = ErrQueueIsEmpty
		return
	}

	err = listQueue.singlyLinkedList.RemoveAtStart()
	if err != nil && errors.Is(err, singlylinkedlist.ErrListIsEmpty) {
		err = ErrQueueIsEmpty
		return
	}
	return
}
