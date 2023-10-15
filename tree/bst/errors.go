package bst

import "errors"

var (
	ErrBSTEmpty = errors.New("tree is empty")
	ErrNotFound = errors.New("not found")
)
