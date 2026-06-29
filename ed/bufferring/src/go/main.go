package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Deque struct {
	data     []int
	front    int
	size     int
	capacity int
}

func (b *Deque) String() string {
	result := []string{}
	for i := range b.size {
		val := b.data[(b.front+i)%b.capacity]
		result = append(result, fmt.Sprint(val))
	}
	return "[" + strings.Join(result, ", ") + "]"
}

func (b *Deque) Debug() string {
	result := make([]string, b.capacity)
	for i, _ := range result {
		result[i] = " _"
		if i == b.front {
			result[i] = ">_"
		}
	}
	for i := range b.size {
		index := (b.front + i) % b.capacity
		val := b.data[index]
		prefix := " "
		if i == 0 {
			prefix = ">"
		}
		result[index] = fmt.Sprintf("%s%d", prefix, val)
	}
	return strings.Join(result, " |")
}

func (b *Deque) Len() int {
	return b.size
}

func (b *Deque) Resize(newCap int) []int{
	newData := make([]int, newCap)
	for i := range b.data {
		newData[i] = b.data[(b.front + i) % b.capacity] 
	}
	b.capacity = newCap
	b.front = 0
	return newData
}

func (b *Deque) PushBack(v int) {
	if b.Len() >= b.capacity {
		b.data = b.Resize(2*b.capacity)
	}
	
	pos := (b.size + b.front) % b.capacity
	b.data[pos] = v
	b.size++
}

func (b *Deque) PushFront(v int) {
	if b.size >= b.capacity {
		b.data = b.Resize(2*b.capacity)
	}
	b.front = (b.front - 1 + b.capacity) % b.capacity
	b.data[b.front] = v
	b.size++
}
func (b *Deque) PopFront(){
	if b.Len() == 0 {
		fmt.Println("fail: buffer vazio")
		return
	}
	b.front = (b.front + 1) % b.capacity
	b.size--
}
func (b *Deque) PopBack(){
	if b.Len() == 0 {
		fmt.Println("fail: buffer vazio")
		return
	}
	b.size--
	
}

func (b *Deque) Front() (int, bool) {
	if b.Len() == 0 {
		return -1, false
	}
	return b.data[b.front], true
}

func (b *Deque) Back() (int, bool){
	if b.Len() == 0 {
		return -1, false
	}
	pos := (b.front + b.size - 1) % b.capacity
	return b.data[pos], true
}

func (b *Deque) Clear(){
	b.size = 0
	b.front = 0
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	buf := &Deque{data: make([]int, 4), capacity: 4}

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
			fmt.Println(buf.String())
		case "debug":
			fmt.Println(buf.Debug())
		case "size":
			fmt.Println(buf.Len())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				buf.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				buf.PushFront(num)
			}
		case "pop_back":
			buf.PopBack();
		case "pop_front":
			buf.PopFront()
		case "front":
			if val, err := buf.Front(); err != true {
				fmt.Println(err)
			} else {
				fmt.Println(val)
			}
		case "back":
			if val, err := buf.Back(); err != true {
				fmt.Println(err)
			} else {
				fmt.Println(val)
			}
		case "clear":
			buf.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
