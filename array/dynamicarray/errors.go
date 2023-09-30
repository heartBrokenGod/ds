package dynamicarray

import "errors"

var (
	ErrInvalidCapacity = errors.New("invalid capacity")
	ErrIndexOutOfRange = errors.New("index out of range")
)
