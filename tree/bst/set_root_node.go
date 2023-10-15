package bst

import "github.com/heartBrokenGod/ds/tree/binarytree"

func (bst *BST[DT]) SetRootNode(rootNode *BSTNode[DT]) {
	bst.bt.RootNode = (*binarytree.BinaryTreeNode[DT])(rootNode)
}
