package staticarray

// Set: set item on the index
//
// If the item already exist on that index,
// the item passed as argument overides the existing value.
//
// This method can return possible errors:
//
// ErrIndexOutOfRange
func (sa *StaticArray[DT]) Set(index int, item DT) error {

	// check if the index is in range
	if err := sa.checkIndexOutOfRange(index); err != nil {
		return err
	}

	if sa.length < index+1 {
		sa.length = index + 1
	}

	// store or override value at index
	sa.arr[index] = item

	// return
	return nil
}
