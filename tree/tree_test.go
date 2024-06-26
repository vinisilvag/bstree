package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBinarySearchTree(t *testing.T) {
	tree := NewBinarySearchTree()
	assert.Nil(t, tree.Root)
	assert.Equal(t, tree.Size, uint(0))
}

func TestInsert(t *testing.T) {
	tree := NewBinarySearchTree()
	tree.insert(10)
	tree.insert(12)
	tree.insert(8)
	tree.insert(9)
	tree.insert(11)
	assert.Equal(t, 10, tree.Root.Elem)
	assert.Equal(t, 12, tree.Root.Right.Elem)
	assert.Equal(t, 8, tree.Root.Left.Elem)
	assert.Equal(t, 9, tree.Root.Left.Right.Elem)
	assert.Equal(t, 11, tree.Root.Right.Left.Elem)
}

func TestSize(t *testing.T) {
	tree := NewBinarySearchTree()
	tree.insert(10)
	tree.insert(12)
	tree.insert(8)
	tree.insert(9)
	tree.insert(11)
	assert.Equal(t, uint(5), tree.Size)
}

func TestMinimum(t *testing.T) {
	tree := NewBinarySearchTree()
	tree.insert(10)
	tree.insert(12)
	tree.insert(8)
	tree.insert(9)
	tree.insert(11)
	minNode := tree.minimum(tree.Root)
	assert.Equal(t, 8, minNode.Elem)
}

func TestMaximum(t *testing.T) {
	tree := NewBinarySearchTree()
	tree.insert(10)
	tree.insert(12)
	tree.insert(8)
	tree.insert(9)
	tree.insert(11)
	maxNode := tree.maximum(tree.Root)
	assert.Equal(t, 12, maxNode.Elem)
}

func TestEmptyTreeMinimum(t *testing.T) {
	tree := NewBinarySearchTree()
	minNode := tree.minimum(tree.Root)
	assert.Nil(t, minNode)
}

func TestEmptyTreeMaximum(t *testing.T) {
	tree := NewBinarySearchTree()
	maxNode := tree.maximum(tree.Root)
	assert.Nil(t, maxNode)
}

func TestSearch(t *testing.T) {
	tree := NewBinarySearchTree()
	tree.insert(10)
	tree.insert(12)
	tree.insert(8)
	tree.insert(9)
	tree.insert(11)
	assert.NotNil(t, tree.search(8))
}

func TestSearchUnexistentElement(t *testing.T) {
	tree := NewBinarySearchTree()
	tree.insert(10)
	tree.insert(12)
	tree.insert(8)
	tree.insert(9)
	tree.insert(11)
	assert.Nil(t, tree.search(13))
}

func TestDeleteLeaf(t *testing.T) {
	tree := NewBinarySearchTree()
	tree.insert(10)
	tree.insert(12)
	tree.insert(8)
	tree.insert(9)
	tree.insert(11)
	tree.insert(13)
	assert.Equal(t, uint(6), tree.Size)
	tree.delete(9)
	assert.Equal(t, uint(5), tree.Size)
	assert.Nil(t, tree.Root.Left.Right)
}

func TestDeleteRootWithOneChild(t *testing.T) {
	tree := NewBinarySearchTree()
	tree.insert(10)
	tree.insert(12)
	tree.insert(8)
	tree.insert(9)
	tree.insert(11)
	assert.Equal(t, uint(5), tree.Size)
	// Case #1
	tree.delete(8)
	assert.Equal(t, uint(4), tree.Size)
	assert.Equal(t, 9, tree.Root.Left.Elem)
	// Case #2
	tree.delete(12)
	assert.Equal(t, uint(3), tree.Size)
	assert.Equal(t, 11, tree.Root.Right.Elem)
}

func TestDeleteSubTreeRoot(t *testing.T) {
	tree := NewBinarySearchTree()
	tree.insert(10)
	tree.insert(12)
	tree.insert(8)
	tree.insert(9)
	tree.insert(11)
	tree.insert(13)
	assert.Equal(t, uint(6), tree.Size)
	tree.delete(12)
	assert.Equal(t, uint(5), tree.Size)
	assert.Equal(t, tree.Root.Right.Elem, 13)
	assert.Equal(t, tree.Root.Right.Left.Elem, 11)
}

func TestDeleteUnexistentElement(t *testing.T) {
	tree := NewBinarySearchTree()
	tree.insert(10)
	tree.insert(12)
	tree.insert(8)
	tree.insert(9)
	tree.insert(11)
	assert.Equal(t, uint(5), tree.Size)
	tree.delete(7)
	assert.Equal(t, uint(5), tree.Size)
}
