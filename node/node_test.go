package node

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNode(t *testing.T) {
	node := NewNode(10)
	assert.Equal(t, 10, node.Elem)
	assert.Nil(t, node.Father)
	assert.Nil(t, node.Left)
	assert.Nil(t, node.Right)
}

func TestNodeIsLeaf(t *testing.T) {
	node := NewNode(10)
	assert.True(t, node.isLeaf())
}

func TestNodeIsNotLeaf(t *testing.T) {
	node := NewNode(10)
	newNode := NewNode(12)
	newNode.Father = &node
	node.Right = &newNode
	assert.False(t, node.isLeaf())
}
