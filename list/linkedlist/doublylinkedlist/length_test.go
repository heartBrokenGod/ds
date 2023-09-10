package doublylinkedlist_test

import (
	"testing"

	"github.com/heartBrokenGod/ds/list/linkedlist/doublylinkedlist"
)

func TestDoublyLinkedList_Length(t *testing.T) {
	dll := doublylinkedlist.New[int]()
	tests := []struct {
		name             string
		doublyLinkedList *doublylinkedlist.DoublyLinkedList[int]
		expectedLength   int
	}{
		{
			name:             "length of list",
			doublyLinkedList: dll,
			expectedLength:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			length := tt.doublyLinkedList.Length()
			if length != tt.expectedLength {
				t.Errorf("DoublyLinkedList.Length() got %d, want %d", length, tt.expectedLength)
			}
		})
	}
}
