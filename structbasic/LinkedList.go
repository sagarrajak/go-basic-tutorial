package structbasic

type Node struct {
	data  float64
	left  *Node
	right *Node
}

func NewNode(value float64, left *Node, right *Node) Node {
	return Node{value, left, right}
}
