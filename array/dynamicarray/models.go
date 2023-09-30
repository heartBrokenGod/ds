package dynamicarray

import "github.com/heartBrokenGod/ds/datatype"

// DynamicArray is an array having fixed length
type DynamicArray[DT datatype.DataType] struct {
	arr      []DT
	length   int
	capacity int
}

func New[DT datatype.DataType](initial_capacity int) (*DynamicArray[DT], error) {
	if initial_capacity <= 0 {
		return nil, ErrInvalidCapacity
	}
	return &DynamicArray[DT]{
		arr:      make([]DT, initial_capacity),
		length:   0,
		capacity: initial_capacity,
	}, nil
}
