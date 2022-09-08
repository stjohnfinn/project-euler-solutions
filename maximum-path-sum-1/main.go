package main

import "fmt"

type Node struct {
	value int
	children []*Node
}

func NewNode(val int) *Node {
	return &Node{
		value: val,
		children: []*Node{},
	}
}

func (n Node) Print() {
	fmt.Printf("Node: %v\n", n.value)
	fmt.Println(n.children)
}

func (n Node) AddChild(child *Node) {
	fmt.Printf("Adding child node: %v", child)
	n.children = append(n.children, child)
}

func main() {
	node1 := NewNode(15)
	node2 := NewNode(46)
	node1.AddChild(node2)
	node1.Print()
}