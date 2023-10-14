package binarytree

import (
	"fmt"

	"github.com/heartBrokenGod/ds/datatype"
	"github.com/heartBrokenGod/ds/queue/listqueue"
)

func preOrder[DT datatype.DataType](btn *BinaryTreeNode[DT]) string {
	s := ""
	if btn != nil {
		s += fmt.Sprint(btn.item, " ")
		s += preOrder(btn.LeftNode)
		s += preOrder(btn.RightNode)
	}
	return s
}

func inOrder[DT datatype.DataType](btn *BinaryTreeNode[DT]) string {
	s := ""
	if btn != nil {
		s += inOrder(btn.LeftNode)
		s += fmt.Sprint(btn.item, " ")
		s += inOrder(btn.RightNode)
	}
	return s
}

func postOrder[DT datatype.DataType](btn *BinaryTreeNode[DT]) string {
	s := ""
	if btn != nil {
		s += postOrder(btn.LeftNode)
		s += postOrder(btn.RightNode)
		s += fmt.Sprint(btn.item, " ")
	}
	return s
}

func levelOrderTraversal[DT datatype.DataType](btn *BinaryTreeNode[DT]) string {
	q := listqueue.New[*BinaryTreeNode[DT]]()
	q.Enqueue(btn)
	s := ""
	for !q.IsEmpty() {
		temp, _ := q.Dequeue()
		s += fmt.Sprint(temp.item, " ")
		if temp.LeftNode != nil {
			q.Enqueue(temp.LeftNode)
		}
		if temp.RightNode != nil {
			q.Enqueue(temp.RightNode)
		}
	}
	return s
}

func (bt *BinaryTree[DT]) String() string {
	switch bt.traversalType {
	case PreOrderTraversal:
		return preOrder(bt.rootNode)
	case InOrderTraversal:
		return inOrder(bt.rootNode)
	case PostOrderTraversal:
		return postOrder(bt.rootNode)
	case LevelOrderTraversal:
		return levelOrderTraversal(bt.rootNode)
	default:
		return preOrder(bt.rootNode)
	}
}
