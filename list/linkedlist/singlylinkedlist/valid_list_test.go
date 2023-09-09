package singlylinkedlist_test

import (
	"testing"

	"github.com/heartBrokenGod/ds/list"
	"github.com/heartBrokenGod/ds/list/linkedlist/singlylinkedlist"
)

func TestListValidity(t *testing.T) {
	var someList list.List[string] = singlylinkedlist.New[string]()
	_ = someList
}
