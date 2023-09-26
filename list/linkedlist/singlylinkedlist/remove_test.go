package singlylinkedlist_test

import (
	"testing"

	"github.com/heartBrokenGod/ds/list/linkedlist/singlylinkedlist"
)

func TestRemove(t *testing.T) {
	sll := singlylinkedlist.New[int]()
	sll.Append(10)
	sll.Append(20)
	sll.Append(30)
	sll.Append(40)
	sll.Append(50)
	type args struct {
		index int
	}

	tests := []struct {
		name             string
		singlyLinkedList *singlylinkedlist.SinglyLinkedList[int]
		args             args
		expectedErr      error
	}{
		{
			name:             "remove at random index",
			singlyLinkedList: sll,
			args: args{
				index: 2,
			},
			expectedErr: nil,
		},
		{
			name:             "remove at start index",
			singlyLinkedList: sll,
			args: args{
				index: 0,
			},
			expectedErr: nil,
		},
		{
			name:             "remove at end index",
			singlyLinkedList: sll,
			args: args{
				index: 2,
			},
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.singlyLinkedList.Remove(tt.args.index)
			if tt.expectedErr != nil && err == tt.expectedErr {
				t.Errorf("SinglyLinkedList.Remove() got err %s, wanted %s", err, tt.expectedErr)
			}
		})
	}

}
