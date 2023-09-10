package doublylinkedlist_test

import (
	"testing"

	"github.com/heartBrokenGod/ds/list/linkedlist/doublylinkedlist"
)

func TestDoublyLinkedList_Prepend(t *testing.T) {
	dll := doublylinkedlist.New[string]()
	type args struct {
		data string
	}
	tests := []struct {
		name             string
		singlyLinkedList *doublylinkedlist.DoublyLinkedList[string]
		args             args
		expectedErr      error
	}{
		{
			name:             "prepend when list is empty",
			singlyLinkedList: dll,
			args: args{
				data: "some string",
			},
			expectedErr: nil,
		},
		{
			name:             "prepend when list is not empty",
			singlyLinkedList: dll,
			args: args{
				data: "some other string",
			},
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.singlyLinkedList.Prepend(tt.args.data)
			if tt.expectedErr != nil && err == tt.expectedErr {
				t.Errorf("DoublyLinkedList.Append() got err %s, wanted %s", err, tt.expectedErr)
			}
		})
	}
}
