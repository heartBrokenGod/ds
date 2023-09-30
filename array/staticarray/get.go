package staticarray

// Get: get item on the index
//
// This method can return possible errors:
//
// ErrIndexOutOfRange
func (sa *StaticArray[DT]) Get(index int) (item DT, err error) {

	// check if the index is in range
	if err := sa.checkIndexOutOfRange(index); err != nil {
		return *new(DT), err
	}

	// return
	return sa.arr[index], nil
}
