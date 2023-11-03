package bst

func (bst *BST[DT]) Insert(item DT, comparator Comparator[DT]) error {

	// check if tree is empty
	if bst.IsEmpty() {
		// add item to root node
		rootNode := NewBSTNode(item)
		bst.SetRootNode(rootNode)
		return nil
	}

	newItemNode := NewBSTNode(item)
	parent := bst.GetRootNode()

	for parent != nil {
		if comparator(parent.Item, newItemNode.Item) > 0 {
			if parent.LeftNode == nil {
				parent.SetLeftNode(newItemNode)
				return nil
			}
			parent = parent.GetLeftNode()
		} else if comparator(parent.Item, newItemNode.Item) < 0 {
			if parent.RightNode == nil {
				parent.SetRightNode(newItemNode)
				return nil
			}
			parent = parent.GetRightNode()
		} else {
			return ErrItemAlreadyExist
		}
	}

	return nil

}
