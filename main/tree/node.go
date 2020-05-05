package tree

import (
	"fmt"
)

type Node struct {
	Value       int
	Left, Right *Node
}

func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	fmt.Println(node.Value)
	node.Right.Traverse()
}
