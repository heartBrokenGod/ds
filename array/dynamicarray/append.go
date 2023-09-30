package dynamicarray

// Append: appends the item at the end of array
//
// If the array length == array capacity,
// then the array will grow increasing the capacity.
func (da *DynamicArray[DT]) Append(item DT) error {

	// check if array needs to grow
	if da.length == da.capacity {
		da.arr = append(da.arr, item)
		da.arr = da.arr[:cap(da.arr)]
		da.length += 1
		da.capacity = cap(da.arr)
	}

	da.arr[da.length] = item
	da.length += 1

	// return
	return nil
}
