package listqueue

import "errors"

var ErrQueueIsEmpty = errors.New("queue is empty")
var ErrQueueIsFull = errors.New("queue is full")
