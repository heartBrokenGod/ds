package dynamicarray

import "fmt"

func (da *DynamicArray[DT]) String() string {
	return fmt.Sprint(da.arr)
}
