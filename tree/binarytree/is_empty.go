package binarytree

func (n *BinaryTree[DT]) IsEmpty() bool {
	return n.RootNode == nil
}
