package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Node[T comparable] struct {
	Value T
	next  *Node[T]
	prev  *Node[T]
	root  *Node[T]
}

type LinkedList[T comparable] struct {
	root *Node[T]
	size int
}

func NewLinkedList[T comparable]() *LinkedList[T] {
	root := &Node[T]{}
	root.next = root
	root.prev = root
	root.root = root
	return &LinkedList[T]{root: root, size: 0}
}

func (n *Node[T]) Next() *Node[T] {
	if n.next == n.root {
		return n.root.next
	}
	return n.next
}
func (n *Node[T]) Prev() *Node[T] {
	if n.prev == n.root {
		return n.root.prev
	}
	return n.prev
}

func (l *LinkedList[T]) PushBack(value T) {
	l.insertBefore(l.root, value)
}

func (l *LinkedList[T]) insertBefore(mark *Node[T], value T) {
	n := &Node[T]{Value: value, root: l.root}
	n.prev = mark.prev
	n.next = mark
	mark.prev.next = n
	mark.prev = n
	l.size++
}


func (l *LinkedList[T]) Front() *Node[T]{
	return l.root.Next()
}

func (l *LinkedList[T]) Back() *Node[T]{
	return l.root.Prev()
}

func (l *LinkedList[T]) Size() int{
	return l.size
}

func (l *LinkedList[T]) Clear(){
	l.root.next = l.root
	l.root.prev = l.root
	l.size = 0
}


func (l *LinkedList[T]) String() string {
	values := []string{}
	for n := l.root.next; n != l.root; n = n.next {
		values = append(values, fmt.Sprint(n.Value))
	}
	return "[" + strings.Join(values, ", ") + "]"
}

func (l *LinkedList[T]) Search(value T) *Node[T] {
	for node := l.Front(); node != nil; node = node.Next(){
		if node.Value == value {
			return node
		}
	}
	return nil
}

func (l *LinkedList[T]) Forward(start T, steps int) bool {
	startNode := l.Search(start)
	if startNode == nil {
		return false
	}
	path := []string{}

	for range steps{
		path = append(path, fmt.Sprintf("%v", startNode.Value))
		startNode = startNode.Next()
	}
	fmt.Printf("[ %s ]\n", strings.Join(path, " "))
	return true
}

func (l *LinkedList[T]) Backward(start T, steps int) bool {
	startNode := l.Search(start)
	if startNode == nil {
		return false
	}
	path := []string{}

	for range steps{
		path = append(path, fmt.Sprintf("%v", startNode.Value))
		startNode = startNode.Prev()
	}
	fmt.Printf("[ %s ]\n", strings.Join(path, " "))
	return true
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewLinkedList[int]()

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			fmt.Println(ll.String())
		case "size":
			fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushBack(num)
			}
		case "clear":
			ll.Clear()
		case "forward":
			start, _ := strconv.Atoi(args[1])
			steps, _ := strconv.Atoi(args[2])

			feedBack := ll.Forward(start, steps)

			if !feedBack {
				fmt.Println("fail: valor não encontrado")
				continue
			}
			// collect := []string{}
			// for range steps {
			// 	collect = append(collect, fmt.Sprintf("%v", node.Value))
			// 	node = node.Next()
			// }
			// fmt.Printf("[ %s ]\n", strings.Join(collect, " "))
		case "backward":
			start, _ := strconv.Atoi(args[1])
			steps, _ := strconv.Atoi(args[2])

			feedBack := ll.Backward(start, steps)

			if !feedBack {
				fmt.Println("fail: valor não encontrado")
				continue
			}
			// collect := []string{}
			// for range steps {
			// 	collect = append(collect, fmt.Sprintf("%v", node.Value))
			// 	node = node.Prev()
			// }
			// fmt.Printf("[ %s ]\n", strings.Join(collect, " "))
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
