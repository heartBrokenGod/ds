package arraystack

func (arrayStack *ArrayStack[DT]) Pop() (item DT, err error) {

	// check if stack is already empty
	if arrayStack.topIndex == -1 {
		item, err = *new(DT), ErrStackUnderFlow
		return
	}

	item = arrayStack.itemArr[arrayStack.topIndex]
	arrayStack.topIndex--
	return

}
