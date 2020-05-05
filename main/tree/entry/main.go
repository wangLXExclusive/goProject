package main

import (
	"main/tree"
)

func main() {
	// var root Node

	// root = Node{value: 3}
	// root.left = &Node{}
	// root.right = &Node{5, nil, nil}
	// root.right.left = new(Node)

	// root.Traverse()
	//var root tree.Node
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Right = &tree.Node{}
	root.Left = &tree.Node{5, nil, nil}
	root.Traverse()
}
