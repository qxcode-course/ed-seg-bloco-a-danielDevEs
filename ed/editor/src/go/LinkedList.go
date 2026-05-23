package main

type LinkedList struct{
    root *Node
	size int
}

func NewList() *LinkedList{
	root := &Node{}
	root.next = root
	root.prev = root
	return &LinkedList{root: root, size: 0}
}

func (l *LinkedList) Front() *Node{
	return l.root.next
}

func (l *LinkedList) Back() *Node{
	return l.root.prev
}

func (l *LinkedList) PushBack(letter rune){
	l.Insert(l.root, letter)
}

func (l *LinkedList) PushFront(letter rune){
	l.Insert(l.root.next, letter)
}

func (l *LinkedList) Erase(Oldnode *Node){
	Oldnode.next.prev = Oldnode.Prev()
	Oldnode.prev.next = Oldnode.Next()
	l.size--
}

func (l *LinkedList) Insert(it *Node, letter rune) *Node{

	prev := it.prev

	newNode := &Node{
		data: letter,
		next: it,
		prev: prev,
	}

	prev.next = newNode
	it.prev = newNode

	return newNode
}

func (l *LinkedList) Replace(node *Node, newValue rune) *Node{
	node.data = newValue
	return node
}