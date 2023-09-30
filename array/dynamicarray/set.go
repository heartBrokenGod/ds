package dynamicarray

// Set: set item on the index
//
// If the item already exist on that index,
// the item passed as argument overides the existing value.
//
// This method can return possible errors:
//
// ErrIndexOutOfRange
func (da *DynamicArray[DT]) Set(index int, item DT) error {

	// check if the index is in range
	if err := da.checkIndexOutOfRange(index); err != nil {
		return err
	}

	if da.length < index+1 {
		da.length = index + 1
	}

	// store or override value at index
	da.arr[index] = item

	// return
	return nil
}
