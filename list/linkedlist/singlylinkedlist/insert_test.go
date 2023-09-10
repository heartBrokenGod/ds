package singlylinkedlist_test

import (
	"testing"

	"github.com/heartBrokenGod/ds/list/linkedlist/singlylinkedlist"
)

func TestSinglyLinkedList_Insert(t *testing.T) {
	sll := singlylinkedlist.New[int]()
	tests := []struct {
		name             string
		singlyLinkedList *singlylinkedlist.SinglyLinkedList[int]
		argIndex         int
		argItem          int
		expectedError    error
	}{
		{
			name:             "insert at begining of list",
			singlyLinkedList: sll,
			argIndex:         0,
			argItem:          100,
			expectedError:    nil,
		},
		{
			name:             "insert at end of list",
			singlyLinkedList: sll,
			argIndex:         1,
			argItem:          200,
			expectedError:    nil,
		},
		{
			name:             "insert at invalid index",
			singlyLinkedList: sll,
			argIndex:         40,
			argItem:          300,
			expectedError:    singlylinkedlist.ErrListIndexOutOfRange,
		},
		{
			name:             "successful insert at valid index",
			singlyLinkedList: sll,
			argIndex:         1,
			argItem:          300,
			expectedError:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.singlyLinkedList.Insert(tt.argIndex, tt.argItem)
			if tt.expectedError != nil && err != tt.expectedError {
				t.Errorf("SinglyLinkedList.Insert() got err %s, want err %s", err, tt.expectedError)
			}
			if err == nil {
				item, _ := tt.singlyLinkedList.Lookup(tt.argIndex)
				if item != tt.argItem {
					t.Errorf("SinglyLinkedList.Insert() could not insert the value")
				}
			}
		})
	}
}
