package binarytree

import (
	"fmt"

	"github.com/heartBrokenGod/ds/datatype"
)

func preOrder[DT datatype.DataType](btn *BinaryTreeNode[DT]) string {
	if btn != nil {
		s := fmt.Sprint(btn.item, ", ")
		s += preOrder(btn.LeftNode)
		s += preOrder(btn.RightNode)
		return s
	}
	return ""
}

func (bt *BinaryTree[DT]) String() string {
	switch bt.traversalType {
	case PreOrderTraversal:
		return preOrder(bt.rootNode)
	default:
		return preOrder(bt.rootNode)
	}
}
