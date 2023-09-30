package staticarray

func (sa *StaticArray[DT]) checkIndexOutOfRange(index int) error {
	if index < 0 || index >= sa.capacity {
		return ErrIndexOutOfRange
	}
	return nil
}
