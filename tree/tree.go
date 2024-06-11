package tree

import (
	"fmt"

	"github.com/vinisilvag/bstree/node"
)

type BinarySearchTree struct {
	Root *node.Node
	Size uint
}

func NewBinarySearchTree() BinarySearchTree {
	return BinarySearchTree{Root: nil, Size: 0}
}

func inorderWalkRecursive(node *node.Node) {
	if node != nil {
		inorderWalkRecursive(node.Left)
		fmt.Println(node.Elem)
		inorderWalkRecursive(node.Right)
	}
}

func (tree BinarySearchTree) inorderWalk() {
	inorderWalkRecursive(tree.Root)
}

func insertRecursive(curr *node.Node, elem int) *node.Node {
	if curr == nil {
		newNode := node.NewNode(elem)
		return &newNode
	} else {
		if elem < curr.Elem {
			curr.Left = insertRecursive(curr.Left, elem)
			curr.Left.Father = curr
		} else {
			curr.Right = insertRecursive(curr.Right, elem)
			curr.Right.Father = curr
		}
	}
	return curr
}

func (tree *BinarySearchTree) insert(elem int) {
	tree.Root = insertRecursive(tree.Root, elem)
	tree.Size += 1
}

func searchRecursive(curr *node.Node, elem int) *node.Node {
	if curr == nil || curr.Elem == elem {
		return curr
	}
	if elem < curr.Elem {
		return searchRecursive(curr.Left, elem)
	} else {
		return searchRecursive(curr.Right, elem)
	}
}

func (tree BinarySearchTree) search(elem int) *node.Node {
	return searchRecursive(tree.Root, elem)
}

func (tree BinarySearchTree) minimum(curr *node.Node) *node.Node {
	if curr == nil {
		return nil
	}
	temp := curr
	for temp.Left != nil {
		temp = temp.Left
	}
	return temp
}

func (tree BinarySearchTree) maximum(curr *node.Node) *node.Node {
	if curr == nil {
		return nil
	}
	temp := curr
	for temp.Right != nil {
		temp = temp.Right
	}
	return temp
}

func transplant(root *node.Node, oldNode *node.Node, newNode *node.Node) {
	if root == oldNode {
		root = newNode
	} else if oldNode == oldNode.Father.Left {
		oldNode.Father.Left = newNode
	} else {
		oldNode.Father.Right = newNode
	}
	if newNode != nil {
		newNode.Father = oldNode.Father
	}
}

func (tree *BinarySearchTree) delete(elem int) {
	nodeToDelete := tree.search(elem)
	if nodeToDelete == nil {
		return
	}
	if nodeToDelete.Left == nil {
		transplant(tree.Root, nodeToDelete, nodeToDelete.Right)
	} else if nodeToDelete.Right == nil {
		transplant(tree.Root, nodeToDelete, nodeToDelete.Left)
	} else {
		successor := tree.minimum(nodeToDelete.Right)
		if successor.Father != nodeToDelete {
			transplant(tree.Root, successor, successor.Right)
			successor.Right = nodeToDelete.Right
			successor.Right.Father = successor
		}
		transplant(tree.Root, nodeToDelete, successor)
		successor.Left = nodeToDelete.Left
		successor.Left.Father = successor
	}
	tree.Size -= 1
}
