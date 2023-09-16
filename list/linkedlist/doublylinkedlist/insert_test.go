package doublylinkedlist_test

import (
	"fmt"
	"testing"

	"github.com/heartBrokenGod/ds/list/linkedlist/doublylinkedlist"
)

func TestDoublyLinkedList_Insert(t *testing.T) {
	dll := doublylinkedlist.New[int]()
	tests := []struct {
		name             string
		doublyLinkedList *doublylinkedlist.DoublyLinkedList[int]
		argIndex         int
		argItem          int
		expectedError    error
	}{
		{
			name:             "insert at begining of list",
			doublyLinkedList: dll,
			argIndex:         0,
			argItem:          100,
			expectedError:    nil,
		},
		{
			name:             "insert at end of list",
			doublyLinkedList: dll,
			argIndex:         1,
			argItem:          200,
			expectedError:    nil,
		},
		{
			name:             "insert at invalid index",
			doublyLinkedList: dll,
			argIndex:         40,
			argItem:          300,
			expectedError:    doublylinkedlist.ErrListIndexOutOfRange,
		},
		{
			name:             "successful insert at valid index closest to tail",
			doublyLinkedList: dll,
			argIndex:         1,
			argItem:          300,
			expectedError:    nil,
		},
		{
			name:             "successful insert at valid index closest to tail 2",
			doublyLinkedList: dll,
			argIndex:         1,
			argItem:          400,
			expectedError:    nil,
		},
		{
			name:             "successful insert at valid index closest to tail 3",
			doublyLinkedList: dll,
			argIndex:         2,
			argItem:          500,
			expectedError:    nil,
		},
		{
			name:             "successful insert at valid index closest to head",
			doublyLinkedList: dll,
			argIndex:         1,
			argItem:          600,
			expectedError:    nil,
		},
		{
			name:             "successful insert at valid index closest to head 2",
			doublyLinkedList: dll,
			argIndex:         2,
			argItem:          700,
			expectedError:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.doublyLinkedList.Insert(tt.argIndex, tt.argItem)
			if tt.expectedError != nil && err != tt.expectedError {
				t.Errorf("DoublyLinkedList.Insert() got err %s, want err %s", err, tt.expectedError)
			}
			if err == nil {
				item, _ := tt.doublyLinkedList.Lookup(tt.argIndex)
				if item != tt.argItem {
					t.Errorf("DoublyLinkedList.Insert() could not insert the value")
				}
			}
			fmt.Println("==========================", tt.name)
			for i := 0; i < dll.Length(); i++ {
				item, _ := dll.Lookup(i)
				fmt.Print(item, ", ")
			}
			fmt.Println("==========================")
		})
	}
}
