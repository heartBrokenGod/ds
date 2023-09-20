package arraystack

func (arrayStack *ArrayStack[DT]) Peek() (item DT, err error) {

	// check if stack is already empty
	if arrayStack.topIndex == -1 {
		item, err = *new(DT), ErrStackEmpty
		return
	}

	item = arrayStack.itemArr[arrayStack.topIndex]
	return

}
