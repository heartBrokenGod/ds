package bst

func (bst *BST[DT]) GetRootNode() *BSTNode[DT] {
	if bst.IsEmpty() {
		return nil
	}
	return (*BSTNode[DT])(bst.bt.RootNode)
}
