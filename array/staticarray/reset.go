package staticarray

func (sa *StaticArray[DT]) Reset(index int) error {

	// check if the index is in range
	if err := sa.checkIndexOutOfRange(index); err != nil {
		return err
	}

	// reset the value at index to zero value of item
	sa.arr[index] = *new(DT)

	// return
	return nil
}
