package dynamicarray

// Get: get item on the index
//
// This method can return possible errors:
//
// ErrIndexOutOfRange
func (da *DynamicArray[DT]) Get(index int) (item DT, err error) {

	// check if the index is in range
	if err := da.checkIndexOutOfRange(index); err != nil {
		return *new(DT), err
	}

	// return
	return da.arr[index], nil
}
