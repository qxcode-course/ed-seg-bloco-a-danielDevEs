package main

type Node struct{
    data rune
    next *Node
    prev *Node
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}
