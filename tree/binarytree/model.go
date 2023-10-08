package binarytree

import "github.com/heartBrokenGod/ds/datatype"

type BinaryTreeNode[DT datatype.DataType] struct {
	item      DT
	LeftNode  *BinaryTreeNode[DT]
	RightNode *BinaryTreeNode[DT]
}

func NewBinaryTreeNode[DT datatype.DataType](item DT) *BinaryTreeNode[DT] {
	return &BinaryTreeNode[DT]{
		item:      item,
		LeftNode:  nil,
		RightNode: nil,
	}
}

type BinaryTree[DT datatype.DataType] struct {
	rootNode      *BinaryTreeNode[DT]
	traversalType TraversalType
}

func New[DT datatype.DataType]() *BinaryTree[DT] {
	return &BinaryTree[DT]{
		rootNode:      nil,
		traversalType: PreOrderTraversal,
	}
}
