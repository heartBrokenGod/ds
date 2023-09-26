package queue

type DataType interface {
	any
}

type Queue[DT DataType] interface {
	Enqueue(item DT) error
	Dequeue() (item DT, err error)
	Peek() (item DT, err error)
}
