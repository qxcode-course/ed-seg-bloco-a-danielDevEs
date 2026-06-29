package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func (q *Queue[T]) Enqueue(value T){
	newNode := &Node[T]{
		Value: value,
	}
	if q.size == 0 {
		q.head = newNode
	} else {
		if q.size == 1 {
			q.tail = newNode
			q.head.next = q.tail
			
		}
		q.tail.next = newNode
		q.tail = newNode
	}
	q.size++
}
func (q *Queue[T]) Dequeue() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}
	if q.size == 1 {
		q.head = nil
		q.size--
		var zero T
		return zero, true
	}
	q.head = q.head.next
	q.size--
	return q.head.Value, true
}
func (q *Queue[T]) Peek() (T, bool){ 
	if q.IsEmpty() {
		var zero T
		return zero, false
	}
	return q.head.Value, true
}
func (q *Queue[T]) Size() int {
	return q.size
}
func (q *Queue[T]) IsEmpty() bool {
	if q.head == nil {
		return true
	}
	return false
}
func (q *Queue[T]) Clear() {
	q.size = 0
	q.head = nil
	q.tail = nil
}


type Node[T any] struct {
	Value T
	next  *Node[T]
}

type Queue[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) String() string {
	result := "["
	for n := q.head; n != nil; n = n.next {
		if n != q.head {
			result += ", "
		}
		result += fmt.Sprintf("%v", n.Value)
	}
	return result + "]"
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	queue := NewQueue[int]()

	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println("$" + line)
		args := strings.Fields(line)
		if len(args) == 0 {
			continue
		}
		switch args[0] {
		case "end":
			break
		case "show":
			fmt.Println(queue)
		case "push":
			for _, arg := range args[1:] {
				value, _ := strconv.Atoi(arg)
				queue.Enqueue(value)
			}
		case "pop":
			if _, ok := queue.Dequeue(); !ok {
				fmt.Println("falha: fila vazia")
			}
		case "peek":
			if value, ok := queue.Peek(); ok {
				fmt.Println(value)
			} else {
				fmt.Println("falha: fila vazia")
			}
		default:
			fmt.Println("Unknown command:", args[0])
		}
	}
}
