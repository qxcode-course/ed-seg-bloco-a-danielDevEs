package main

import (
	"bufio"
	"fmt"
	"os"
	// "slices"
	"strconv"
	"strings"
)

// func partition (vet []int, low, high int) int {
// 	pivot := vet[high]
// 	i := low - 1

// 	for j := low; j < high; j++{
// 		if vet[j] < pivot{
// 			i++
// 			vet[i], vet[j] = vet[j], vet[i]
// 		}
// 		vet[i+1], vet[high] = vet[high], vet[i+1]

// 		return i + 1
// 	}
// 	return -1
// }
// func quickSort(vet []int, low, high int) []int{
// 	if low < high {
// 		pivot := partition(vet, low, high)

// 		quickSort(vet, low, pivot-1)
// 		quickSort(vet, pivot+1, high)
// 	}
// 	return vet
// }

// func quickSort(arr []int) []int {
// 	if len(arr) < 2 {
// 		return arr
// 	}

// 	pivo := arr[0]

// 	var menores []int
// 	var maiores []int

// 	for _, v := range arr[1:] {
// 		if v <= pivo {
// 			menores = append(menores, v)
// 		} else {
// 			maiores = append(maiores, v)
// 		}
// 	}

// 	menores = quickSort(menores)
// 	maiores = quickSort(maiores)

// 	resultado := append(menores, pivo)
// 	resultado = append(resultado, maiores...)

// 	return resultado
// }

type Set struct{
	data		[]int
	size		int
	capacity	int
}

func NewSet(capacity int) *Set{
	return &Set{
		data: 		make([]int, capacity),
		size: 		0,
		capacity: 	capacity,
	}
}

func (s *Set) Insert(value int){
	if s.Contains(value){
		return
	}
	if s.size == s.capacity{
		s.Reserve(s.capacity * 2)
		newData := NewSet(s.capacity)
		for i:= 0; i < s.size; i++{
			newData.data[i] = s.data[i]
		}
		s.data = newData.data
	}

	pos := 0

	for pos < s.size && s.data[pos] < value{
		pos++
	}

	for i := s.size; i > pos; i--{
		s.data[i] = s.data[i-1]
	}




	s.data[pos] = value
	s.size++
}

func (s *Set) Show() string{
	// slices.Sort(s.data)
	if s.size == 0{
		return "[]"
	}
	str := ""
	str += "["
	for i := 0; i < s.size; i++{
		if s.data[i] == 0{
			continue
		}
		str += fmt.Sprint(s.data[i])
		if i < s.size - 1{
			str += ", "
		}
	}
	str += "]"
	return str
}

func (s *Set) reserve(newCapacity int){
	s.capacity = newCapacity
}

func (s *Set) binarySearch(value int) int{
	inicio := 0
	fim := s.size - 1
	meio := inicio+fim/2
	if s.data[meio] == value{
		return meio
	} else if s.data[meio] < value{
		inicio = meio + 1 
	} else if s.data[meio] > value{
		fim = meio - 1
	}
	for true {
		if inicio > fim{
			break
		}
		if s.data[inicio] == value{
			return inicio
		}
		inicio++
	}
	return -1
}

func (s *Set) Erase(value int) bool{
	if !s.Contains(value){
		return false
	}
	pos := s.binarySearch(value)
	for i := pos; i < s.size - 1; i++{
		s.data[i] = s.data[i+1]
	}
	s.size--
	return true
}


func (s *Set) Contains(value int) bool{
	if s.binarySearch(value) != -1{
		return true
	}
	return false
}

func (s *Set) Reserve(newCapacity int){
	s.capacity = newCapacity
}


func main() {

	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewSet(0)
	for scanner.Scan() {
		fmt.Print("$")
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
			v = NewSet(value)
		case "insert":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.Insert(value)
			}
		case "show":
			fmt.Println(v.Show())
		case "erase":
			value, _ := strconv.Atoi(parts[1])
			err := v.Erase(value)
			if err == false{
				fmt.Println("value not found")
			}
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			fmt.Println(v.Contains(value))
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
