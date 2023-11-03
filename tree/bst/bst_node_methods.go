package bst

import "github.com/heartBrokenGod/ds/tree/binarytree"

func (n *BSTNode[DT]) GetLeftNode() *BSTNode[DT] {
	return (*BSTNode[DT])(n.LeftNode)
}

func (n *BSTNode[DT]) GetRightNode() *BSTNode[DT] {
	return (*BSTNode[DT])(n.RightNode)
}

func (n *BSTNode[DT]) SetLeftNode(node *BSTNode[DT]) {
	n.LeftNode = (*binarytree.BinaryTreeNode[DT])(node)
}

func (n *BSTNode[DT]) SetRightNode(node *BSTNode[DT]) {
	n.RightNode = (*binarytree.BinaryTreeNode[DT])(node)
}
