package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)
type linkedDlist struct{
	root *Node
	size int
}

func NewLList() *linkedDlist{
	root := &Node{}
	root.next = root
	root.prev = root
	root.root = root
	return &linkedDlist{root: root, size: 0}
}

func (l *linkedDlist) String() string {
	str := ""
	if l.size != 0{
		str += "["
		for node := l.Front(); node != nil; node = node.Next() {
			str += fmt.Sprintf("%d", node.value)

			if node.next != l.root {
					str += ", "
				}
		}
		
		str += "]"
		return str
	}
	str += "[ ]"
	return str
}

func (l *linkedDlist) Size() int {
	return l.size
}

func (l *linkedDlist) Clear(){
	l.root.next = l.root
	l.root.prev = l.root
	l.size = 0
}

func (l *linkedDlist) PushFront(value int){
	l.Insert(l.root.next, value)
}

func (l *linkedDlist) PushBack(value int){
	l.Insert(l.root, value)
}

func (l *linkedDlist) Insert(it *Node, value int) *Node{
	newNode := &Node{
		value: value,
		root: l.root,
	}

	// Insere no meio
	newNode.next = it
	newNode.prev = it.prev

	// Faz a costura
	it.prev.next = newNode
	it.prev = newNode

	l.size++
	return newNode

}

func (l *linkedDlist) Front() *Node{
	if l.root.next == l.root{
		return nil
	}
	return l.root.next
}

func (l *linkedDlist) Back() *Node{
	if l.root.prev == l.root{
		return nil
	}
	return l.root.prev
}

func (l *linkedDlist) Search(value int) *Node {
	for node := l.Front(); node != nil; node = node.Next() {
		if node.value == value {
			return node
		}
	}
	return nil
}

func (l *linkedDlist) Remove(value int) bool{
	node := l.Search(value)
	if node == nil {
		return false
	}
	node.prev.next = node.Next()
	node.next.prev = node.Prev()
	l.size--
	return true
}

func (l *linkedDlist) Replace(oldvalue, newvalue int) bool {
	node := l.Search(oldvalue)
	if node == nil {
		return false
	}
	node.value = newvalue
	return true
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewLList()

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
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			}
		case "pop_back":
			// ll.PopBack()
		case "pop_front":
			// ll.PopFront()
		case "clear":
			ll.Clear()
		case "walk":
			fmt.Print("[ ")
			for node := ll.Front(); node != nil; node = node.Next() {
				fmt.Printf("%v ", node.value)
			}
			fmt.Print("]\n[ ")
			for node := ll.Back(); node != nil; node = node.Prev() {
				fmt.Printf("%v ", node.value)
			}
			fmt.Println("]")
		case "replace":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			
			feedback := ll.Replace(oldvalue, newvalue)
			if feedback == false{
				fmt.Println("fail: not found")
			}
		case "insert":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Insert(node, newvalue)
			} else {
				fmt.Println("fail: not found")
			}
		case "remove":
			oldvalue, _ := strconv.Atoi(args[1])
			feedback := ll.Remove(oldvalue)
			if feedback == false {
				fmt.Println("fail: not found")
			}
			// node := ll.Search(oldvalue)
			// if node != nil {
			// } else {
			// 	fmt.Println("fail: not found")
			// }
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
