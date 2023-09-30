package staticarray

import "github.com/heartBrokenGod/ds/datatype"

// StaticArray is an array having fixed length
type StaticArray[DT datatype.DataType] struct {
	arr      []DT
	length   int
	capacity int
}

func New[DT datatype.DataType](capacity int) (*StaticArray[DT], error) {
	if capacity <= 0 {
		return nil, ErrInvalidCapacity
	}
	return &StaticArray[DT]{
		arr:      make([]DT, capacity),
		length:   0,
		capacity: capacity,
	}, nil
}
