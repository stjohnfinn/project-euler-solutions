package main

import "fmt"

const WIDTH int = 2
const HEIGHT int = 2

type Node struct {
	x          int
	y          int
	canGoDown  bool
	canGoRight bool
}

func createNode(x int, y int) *Node {
	var out *Node = new(Node)
	out.x = x
	out.y = y
	if y == HEIGHT {
		out.canGoDown = false
	} else {
		out.canGoDown = true
	}
	if x == WIDTH {
		out.canGoRight = false
	} else {
		out.canGoDown = true
	}
	return out
}

func main() {
	var n *Node = createNode(0, 0)
	fmt.Println(n.canGoDown)
}
