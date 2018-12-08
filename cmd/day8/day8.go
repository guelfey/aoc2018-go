package main

import "fmt"

type Node struct {
	children []Node
	metadata []int
}

func readInt() int {
	var i int
	_, err := fmt.Scan(&i)
	if err != nil {
		panic(err)
	}
	return i
}

func readNode() Node {
	nchildren := readInt()
	nmetadata := readInt()
	n := Node{
		children: make([]Node, nchildren),
		metadata: make([]int, nmetadata),
	}
	for i := range n.children {
		n.children[i] = readNode()
	}
	for i := range n.metadata {
		n.metadata[i] = readInt()
	}
	return n
}

func (n Node) metadataSum() int {
	s := 0
	for _, i := range n.metadata {
		s += i
	}
	for _, c := range n.children {
		s += c.metadataSum()
	}
	return s
}

func (n Node) value() int {
	if len(n.children) == 0 {
		return n.metadataSum()
	}
	v := 0
	for _, idx := range n.metadata {
		if idx <= len(n.children) {
			v += n.children[idx-1].value()
		}
	}
	return v
}

func main() {
	root := readNode()
	fmt.Println(root.metadataSum())
	fmt.Println(root.value())
}
