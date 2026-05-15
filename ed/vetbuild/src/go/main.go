package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vector struct {
	data     []int
	size     int
	capacity int
}

func NewVector(capacity int) *Vector {
	return &Vector{
		data:     make([]int, capacity), // nunca use len(data) ou cap(data) ou qq método do go de manipulação de array
		size:     0,
		capacity: capacity,
	}
}

func (v *Vector)Status() string{
	str := ""
	str += fmt.Sprintf("size:%d capacity:%d", v.size, v.capacity)
	return str
}
func (v *Vector)String() string{
	str := "["
	for i := 0; i < v.size; i++{
		if v.data[i] != 0{
			str+= fmt.Sprint(v.data[i])
		}
		if i < v.size -1{
			str += ", "
		}
	}
	// str += fmt.Sprintf("size:%d capacity:%d", v.size, v.capacity)
	str += "]"
	return str
}
func (v *Vector) Capacity() int {
	return v.capacity
} 
func (v *Vector) Size() int {
	return v.size
} 

func (v *Vector) PushBack(value int){
	if(v.size == v.capacity){
		v.capacity*=2
		newData := make([]int, v.capacity)
		for i := 0; i < v.size; i++{
			newData[i] = v.data[i]
		}
		v.data = newData
	}
	v.data[v.size] = value
	v.size++
}

func (v *Vector) Get(pos int) int{
	if pos < 0 || pos >= v.size{
		return -1
	}
	return v.data[pos]
}

func (v *Vector) Set(pos, value int) bool{
	if pos < 0 || pos >= v.size{
		return false
	}
	v.data[pos] = value
	return true
}

func (v *Vector) Clear(){
	v.size = 0
}

func (v *Vector) Reserve(value int){
	v.capacity = value
}

func (v *Vector) PopBack() []int{
	if(v.size == 0){
		return nil
	}
	v.size--
	return v.data
}

func (v *Vector) Insert(pos, value int)[]int{
	if pos < 0 || pos >= v.size{
		return nil
	}
	if(v.size == v.capacity){
		v.capacity*=2

		newData := make([]int, v.capacity)
		for i := 0; i < v.size; i++{
			newData[i] = v.data[i]
		}
		v.data = newData
	}
	
	salvar := make([]int, v.size)
	for i := 0; i < v.size - pos; i++ {
		salvar[i] = v.data[pos+i]
	}
	v.data[pos] = value
	v.size++
	for i := 0; i < v.size - pos - 1; i++ {
		v.data[pos + 1 + i] = salvar[i]
	}
	return v.data
}

func (v *Vector) Erase(pos int) []int{
	if pos < 0 || pos >= v.size{
		return nil
	}
	for i := 0; i < v.size-pos-1; i++{
		v.data[i + pos] = v.data[pos + i + 1]
	}
	v.size--
	return v.data
}

func (v *Vector) IndexOf(value int) int{
	for i := 0; i < v.size; i++{
		if v.data[i] == value{
			return i
		}
	}
	return -1
}

func (v *Vector) Contains(value int) bool{
	for i := 0; i < v.size; i++{
		if v.data[i] == value{
			return true
		}
	}
	return false
}

func (v *Vector) Slice(posI, posF int) []int{
	if posF == -1{
		posF = v.size - 1
	}
	slice := make([]int, posF - posI)

	for i := 0; i < posF - posI; i++{
		slice[i] = v.data[posI + i]
	}
	return slice
}


func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return "[]"
	}
	var result strings.Builder
	fmt.Fprintf(&result, "[%d", slice[0])
	for _, value := range slice[1:] {
		fmt.Fprintf(&result, "%s%d", sep, value)
	}
	fmt.Fprintf(&result, "]")
	return result.String()
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	

	v := NewVector(0)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			v = NewVector(value)
		case "push":
			for _, part := range parts[1:] {
			 	value, _ := strconv.Atoi(part)
			 	v.PushBack(value)
			}
		case "show":
			fmt.Println(v)
		case "status":
			fmt.Println(v.Status())
		case "pop":
			err := v.PopBack()
			if err == nil {
				fmt.Println("vector is empty")
			}
		case "insert":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Insert(index, value)
			if err == nil {
				fmt.Println("index out of range")
			}
		case "erase":
			index, _ := strconv.Atoi(parts[1])
			err := v.Erase(index)
			if err == nil {
				fmt.Println("index out of range")
			}
		case "indexOf":
			value, _ := strconv.Atoi(parts[1])
			index := v.IndexOf(value)
			// if index == -1{
			// 	fmt.Println("Ivalid value")
			// }
			fmt.Println(index)
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			if v.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "clear":
			v.Clear()
		case "capacity":
			fmt.Println(v.Capacity())
		case "get":
			index, _ := strconv.Atoi(parts[1])
			value := v.Get(index)
			if value == -1 {
				fmt.Println("index out of range")
			} else {
				fmt.Println(value)
			}
		case "set":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Set(index, value)
			if err == false {
				fmt.Println("index out of range")
			}
			
		case "reserve":
			newCapacity, _ := strconv.Atoi(parts[1])
			v.Reserve(newCapacity)
		case "slice":
			start, _ := strconv.Atoi(parts[1])
			end, _ := strconv.Atoi(parts[2])
			slice := v.Slice(start, end)
			fmt.Println(Join(slice, ", "))
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
