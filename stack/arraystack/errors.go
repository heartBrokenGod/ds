package arraystack

import "errors"

var (
	ErrInvalidStackSize = errors.New("invalid stack size")
	ErrStackOverFlow    = errors.New("stack overflow")
	ErrStackUnderFlow   = errors.New("stack underflow")
	ErrStackEmpty       = errors.New("stack is empty")
)
