package node

type Node struct {
	Elem   int
	Father *Node
	Left   *Node
	Right  *Node
}

func NewNode(elem int) Node {
	return Node{Elem: elem, Father: nil, Left: nil, Right: nil}
}

func (node Node) isLeaf() bool {
	return node.Left == nil && node.Right == nil
}
