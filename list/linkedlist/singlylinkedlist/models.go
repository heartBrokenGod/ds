package singlylinkedlist

type DataType interface {
	any
}

type Node[DT DataType] struct {
	Data DT
	Next *Node[DT]
}

type SinglyLinkedList[DT DataType] struct {
	length int
	head   *Node[DT]
	tail   *Node[DT]
}

func New[DT DataType]() *SinglyLinkedList[DT] {
	return &SinglyLinkedList[DT]{
		length: 0,
		head:   nil,
		tail:   nil,
	}
}
