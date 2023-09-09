package doublylinkedlist

type DataType interface {
	any
}

type Node[DT DataType] struct {
	Data DT
	Prev *Node[DT]
	Next *Node[DT]
}

type DoublyLinkedList[DT DataType] struct {
	length int
	head   *Node[DT]
	tail   *Node[DT]
}

func New[DT DataType]() *DoublyLinkedList[DT] {
	return &DoublyLinkedList[DT]{
		length: 0,
		head:   nil,
		tail:   nil,
	}
}
