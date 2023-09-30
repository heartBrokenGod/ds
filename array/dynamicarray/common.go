package dynamicarray

func (da *DynamicArray[DT]) checkIndexOutOfRange(index int) error {
	if index < 0 || index >= da.capacity {
		return ErrIndexOutOfRange
	}
	return nil
}
