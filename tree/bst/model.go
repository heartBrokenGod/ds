package bst

import (
	"github.com/heartBrokenGod/ds/datatype"
	"github.com/heartBrokenGod/ds/tree/binarytree"
)

type BSTNode[DT datatype.DataType] binarytree.BinaryTreeNode[DT]

type BST[DT datatype.DataType] struct {
	bt *binarytree.BinaryTree[DT]
}

func New[DT datatype.DataType]() *BST[DT] {
	return &BST[DT]{
		bt: binarytree.New[DT](),
	}
}
