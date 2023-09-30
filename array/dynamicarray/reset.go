package dynamicarray

func (da *DynamicArray[DT]) Reset(index int) error {

	// check if the index is in range
	if err := da.checkIndexOutOfRange(index); err != nil {
		return err
	}

	// reset the value at index to zero value of item
	da.arr[index] = *new(DT)

	// return
	return nil
}
