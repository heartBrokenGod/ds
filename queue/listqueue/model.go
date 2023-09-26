package listqueue

import (
	"github.com/heartBrokenGod/ds/list/linkedlist/singlylinkedlist"
)

type DataType interface {
	any
}

type ListQueue[DT DataType] struct {
	singlyLinkedList *singlylinkedlist.SinglyLinkedList[DT]
}

func New[DT DataType]() *ListQueue[DT] {
	return &ListQueue[DT]{
		singlyLinkedList: singlylinkedlist.New[DT](),
	}
}
