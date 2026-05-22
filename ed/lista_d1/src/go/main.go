package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)



type Node struct{
	data int
	next *Node
	prev *Node
}


type linkedList struct{
	root *Node
}

func newList()*linkedList{
    l := linkedList{}

    l.root = &Node{}
    l.root.next = l.root
    l.root.prev = l.root

    return &l
}

func (l *linkedList) Show() string{
	str := ""
	str += "["
	for i := l.root.next; i != l.root; i = i.next {
		str += fmt.Sprintf("%d", i.data)

		if i.next != l.root {
				str += ", "
			}
	}
	
	str += "]"
	return str
}

func (l *linkedList) Size() int{
	size := 0
	for i := l.root.next; i != l.root; i = i.next {
		size++
	}
	return size
}

func (l *linkedList) PushBack(value int) {
	newNode := &Node{
		data: value,
	}

	newNode.next = l.root
	newNode.prev = l.root.prev

	l.root.prev.next = newNode
	l.root.prev = newNode
}

func (l *linkedList) PushFront(value int) {
	newNode := &Node{
			data: value,
		}

	newNode.next = l.root.next
	newNode.prev = l.root

	l.root.next.prev = newNode
	l.root.next = newNode
}

// func inserir(node *Node, value int){ // COMO IMPLEMENTA?

// }

func (l *linkedList) PopBack(){
	l.root.prev = l.root.prev.prev
	l.root.prev.next = l.root
}

func (l *linkedList) PopFront(){
	l.root.next = l.root.next.next
	l.root.next.prev = l.root
}

func (l *linkedList) Clear(){
	l.root.next = l.root
	l.root.prev = l.root

}



// func inserir(node *Node, value int){
//     novo := &Node{
//         data: value,
//         next: node,
//         prev: node.prev,
//     }
// 	node.prev.next: novo,
// 	node.prev: novo,
// }



func main() {
	scanner := bufio.NewScanner(os.Stdin)
 	l := newList()

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
			fmt.Println(l.Show())
		case "size":
			fmt.Println(l.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				l.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				l.PushFront(num)
			}
		case "pop_back":
			l.PopBack()
		case "pop_front":
			l.PopFront()
		case "clear":
			l.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading input:", err)
	}
}
