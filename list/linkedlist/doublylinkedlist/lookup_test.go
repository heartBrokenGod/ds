package doublylinkedlist_test

import (
	"testing"

	"github.com/heartBrokenGod/ds/list/linkedlist/doublylinkedlist"
)

func TestDoublyLinkedList_Lookup(t *testing.T) {
	emptydll := doublylinkedlist.New[int]()
	dll := doublylinkedlist.New[int]()
	dll.Append(100)
	dll.Append(200)
	dll.Append(300)
	dll.Append(400)
	tests := []struct {
		name             string
		doublyLinkedList *doublylinkedlist.DoublyLinkedList[int]
		argIndex         int
		expectedItem     int
		expectedError    error
	}{
		{
			name:             "list empty",
			doublyLinkedList: emptydll,
			expectedError:    doublylinkedlist.ErrListIsEmpty,
		},
		{
			name:             "index out of range",
			doublyLinkedList: dll,
			argIndex:         4,
			expectedItem:     0,
			expectedError:    doublylinkedlist.ErrListIndexOutOfRange,
		},
		{
			name:             "successful lookup index closest to head",
			doublyLinkedList: dll,
			argIndex:         1,
			expectedItem:     200,
			expectedError:    nil,
		},
		{
			name:             "successful lookup index closest to tail",
			doublyLinkedList: dll,
			argIndex:         2,
			expectedItem:     300,
			expectedError:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item, err := tt.doublyLinkedList.Lookup(tt.argIndex)
			if item != tt.expectedItem {
				t.Errorf("DoublyLinkedList.Lookup() got item %d, want item %d", item, tt.expectedItem)
			}
			if tt.expectedError != nil && err != tt.expectedError {
				t.Errorf("DoublyLinkedList.Lookup() got err %s, want err %s", err, tt.expectedError)
			}
		})
	}
}
