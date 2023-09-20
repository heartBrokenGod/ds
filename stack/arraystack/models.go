package arraystack

type DataType interface {
	any
}

type ArrayStack[DT DataType] struct {
	itemArr     []DT
	stackHeight int
	topIndex    int
}

func New[DT DataType](stackHeight int) (*ArrayStack[DT], error) {

	// check for invalid stack size
	if stackHeight < 1 {
		return nil, ErrInvalidStackSize
	}

	return &ArrayStack[DT]{
		itemArr:     make([]DT, stackHeight),
		stackHeight: stackHeight,
		topIndex:    -1,
	}, nil

}
