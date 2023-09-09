package singlylinkedlist_test

import (
	"testing"

	"github.com/heartBrokenGod/ds/list/linkedlist/singlylinkedlist"
)

func TestSinglyLinkedList_Length(t *testing.T) {
	sll := singlylinkedlist.New[int]()
	tests := []struct {
		name             string
		singlyLinkedList *singlylinkedlist.SinglyLinkedList[int]
		expectedLength   int
	}{
		{
			name:             "length of list",
			singlyLinkedList: sll,
			expectedLength:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			length := tt.singlyLinkedList.Length()
			if length != tt.expectedLength {
				t.Errorf("SinglyLinkedList.Length() got %d, want %d", length, tt.expectedLength)
			}
		})
	}
}
