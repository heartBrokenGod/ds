package bst

import (
	"github.com/heartBrokenGod/ds/datatype"
	"github.com/heartBrokenGod/ds/tree/binarytree"
)

type BSTNode[DT datatype.DataType] binarytree.BinaryTreeNode[DT]

func NewBSTNode[DT datatype.DataType](item DT) *BSTNode[DT] {
	return &BSTNode[DT]{
		Item:      item,
		LeftNode:  nil,
		RightNode: nil,
	}
}

type BST[DT datatype.DataType] struct {
	bt *binarytree.BinaryTree[DT]
}

func New[DT datatype.DataType]() *BST[DT] {
	return &BST[DT]{
		bt: binarytree.New[DT](),
	}
}
