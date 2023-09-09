package singlylinkedlist

import "errors"

var (
	errPrefix              = "singly linked list: "
	ErrListIsEmpty         = errors.New(errPrefix + "list is empty")
	ErrListIndexOutOfRange = errors.New(errPrefix + "index is out of list range")
)
