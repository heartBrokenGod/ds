package binarytree

import "github.com/heartBrokenGod/ds/datatype"

type BinaryTreeNode[DT datatype.DataType] struct {
	Item      DT
	LeftNode  *BinaryTreeNode[DT]
	RightNode *BinaryTreeNode[DT]
}

func NewBinaryTreeNode[DT datatype.DataType](item DT) *BinaryTreeNode[DT] {
	return &BinaryTreeNode[DT]{
		Item:      item,
		LeftNode:  nil,
		RightNode: nil,
	}
}

type BinaryTree[DT datatype.DataType] struct {
	RootNode      *BinaryTreeNode[DT]
	TraversalType TraversalType
}

func New[DT datatype.DataType]() *BinaryTree[DT] {
	return &BinaryTree[DT]{
		RootNode:      nil,
		TraversalType: PreOrderTraversal,
	}
}
