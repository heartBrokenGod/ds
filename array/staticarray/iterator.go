package staticarray

import (
	"github.com/heartBrokenGod/ds/datatype"
)

type Iterator[DT datatype.DataType] interface {
	HasNext() bool
	HasPrev() bool
	Next() DT
	Prev() DT
	Begin() Iterator[DT]
	End() Iterator[DT]
}

type IteratorImpl[DT datatype.DataType] struct {
	sa     *StaticArray[DT]
	cursor int // index
}

func (i *IteratorImpl[DT]) HasNext() bool {
	return i.cursor+1 < i.sa.capacity
}

func (i *IteratorImpl[DT]) HasPrev() bool {
	return i.cursor-1 > -1
}

func (i *IteratorImpl[DT]) Next() DT {
	if i.cursor+1 >= i.sa.capacity {
		panic(ErrIndexOutOfRange)
	}
	i.cursor++
	item, _ := i.sa.Get(i.cursor)
	return item
}

func (i *IteratorImpl[DT]) Prev() DT {
	if i.cursor-1 <= -1 {
		panic(ErrIndexOutOfRange)
	}
	i.cursor--
	item, _ := i.sa.Get(i.cursor)
	return item
}

func (i *IteratorImpl[DT]) Begin() Iterator[DT] {
	i.cursor = -1
	return i
}

func (i *IteratorImpl[DT]) End() Iterator[DT] {
	i.cursor = i.sa.capacity
	return i
}

func (sa *StaticArray[DT]) Iterator() Iterator[DT] {

	return &IteratorImpl[DT]{
		cursor: -1,
		sa:     sa,
	}

}
