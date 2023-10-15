package bst

import "github.com/heartBrokenGod/ds/datatype"

// return negative number if item1 is less than item2
// return positive number if item1 is greater than item2
// return zero if both items are equal
type Comparator[DT datatype.DataType] func(item1, item2 DT) int

func findNode[DT datatype.DataType](bstn *BSTNode[DT], item DT, comparator Comparator[DT]) *BSTNode[DT] {
	if bstn == nil {
		return nil
	}
	compResult := comparator(item, bstn.Item)
	if compResult < 0 {
		return findNode((*BSTNode[DT])(bstn.LeftNode), item, comparator)
	} else if compResult > 0 {
		return findNode((*BSTNode[DT])(bstn.RightNode), item, comparator)
	} else {
		return bstn
	}

}

func (bst *BST[DT]) Find(item DT, comparator Comparator[DT]) (*BSTNode[DT], error) {
	if bst.IsEmpty() {
		return nil, ErrBSTEmpty
	}
	bstn := findNode[DT]((*BSTNode[DT])(bst.bt.RootNode), item, comparator)
	if bstn != nil {
		return bstn, nil
	}
	return nil, ErrNotFound
}
