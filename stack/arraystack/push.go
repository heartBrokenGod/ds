package arraystack

func (arrayStack *ArrayStack[DT]) Push(item DT) error {

	// check if stack is already full
	if arrayStack.stackHeight-1 == arrayStack.topIndex {
		return ErrStackOverFlow
	}

	arrayStack.topIndex++
	arrayStack.itemArr[arrayStack.topIndex] = item
	return nil

}
