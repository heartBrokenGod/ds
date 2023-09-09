package list

import (
	"github.com/heartBrokenGod/ds/list/linkedlist/doublylinkedlist"
	"github.com/heartBrokenGod/ds/list/linkedlist/singlylinkedlist"
)

type DataType interface {
	singlylinkedlist.DataType | doublylinkedlist.DataType
}

type List[DT DataType] interface {
	Length() int
	Prepend(item DT) error
	Append(item DT) error
}

type ListImplType int

const (
	SinglyLinkedList ListImplType = iota
	DoublyLinkedList
	CircularLinkedList
)

func New[DT DataType](listType ListImplType) (ls List[DT], err error) {
	switch listType {
	case SinglyLinkedList:
		ls = singlylinkedlist.New[DT]()
	case DoublyLinkedList:
		ls = doublylinkedlist.New[DT]()
	default:
		err = ErrListImplementationTypeNotImplemented
	}
	return

}
