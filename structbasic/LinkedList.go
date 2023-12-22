package structbasic

type Node struct {
	Data  float64
	Left  *Node
	Right *Node
}

func NewNode(value float64, left *Node, right *Node) *Node {
	return &Node{value, left, right}
}
