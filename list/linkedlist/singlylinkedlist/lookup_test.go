package singlylinkedlist_test

import (
	"testing"

	"github.com/heartBrokenGod/ds/list/linkedlist/singlylinkedlist"
)

func TestSinglyLinkedList_Lookup(t *testing.T) {
	emptysll := singlylinkedlist.New[int]()
	sll := singlylinkedlist.New[int]()
	sll.Append(100)
	sll.Append(200)
	tests := []struct {
		name             string
		singlyLinkedList *singlylinkedlist.SinglyLinkedList[int]
		argIndex         int
		expectedItem     int
		expectedError    error
	}{
		{
			name:             "list empty",
			singlyLinkedList: emptysll,
			expectedError:    singlylinkedlist.ErrListIsEmpty,
		},
		{
			name:             "index out of range",
			singlyLinkedList: sll,
			argIndex:         4,
			expectedItem:     0,
			expectedError:    singlylinkedlist.ErrListIndexOutOfRange,
		},
		{
			name:             "successful lookup",
			singlyLinkedList: sll,
			argIndex:         1,
			expectedItem:     200,
			expectedError:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item, err := tt.singlyLinkedList.Lookup(tt.argIndex)
			if item != tt.expectedItem {
				t.Errorf("SinglyLinkedList.Lookup() got item %d, want item %d", item, tt.expectedItem)
			}
			if tt.expectedError != nil && err != tt.expectedError {
				t.Errorf("SinglyLinkedList.Lookup() got err %s, want err %s", err, tt.expectedError)
			}
		})
	}
}
