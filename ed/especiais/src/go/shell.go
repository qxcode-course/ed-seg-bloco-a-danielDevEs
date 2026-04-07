package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	One int
	Two int
}


func abs(n int) int {
	if n < 0{
		return -n
	}
	return n
}

func quickSort(pares []Pair, low, high int) []Pair{
	if low < high{
		p := partition(pares, low, high)

		quickSort(pares, low, p-1)
		quickSort(pares, p+1, high)
	}
	return pares
}

func partition(pares []Pair, low, high int) int{
	pivot := pares[high]
	i := low - 1
	
	for j := low; j < high; j++{
		if pares[j].One < pivot.One{
			i++
			pares[i].One, pares[j].One = pares[j].One, pares[i].One
		}
	}


	pares[i+1], pares[high] = pares[high], pares[i+1]
	return i+1
}

func occurr(vet []int) []Pair {
	freq := make(map[int]int)
	var ocur []Pair

	for _, pos := range vet{
		freq[abs(pos)]++
		//freq[abs(vet[pos])] = true
	}
	//fmt.Println(freq)
	for elemento := range freq{
		ocur = append(ocur, Pair{One: elemento, Two: freq[elemento]})
	}
	
	return quickSort(ocur, 0, len(freq) - 1)
	
}

func teams(vet []int) []Pair {
	freq := make(map[int]int)
	var teams []Pair
	// unic := 1
	// for v := 1; v < len(vet); v++{
	// 	if vet[v - 1] == vet[v] && len(vet) > 1{
	// 		unic++
	// 		teams = append(teams, Pair{One: vet[v - 1], Two: unic})
	// 	} else{
	// 		teams = append(teams, Pair{One: vet[v - 1], Two: 1})
	// 		unic = 1
	// 	}
	// }
	unics := make(map[int]bool)
	count := 1
	storageCount := []int{}
	for i := 1; i < len(vet); i++{
		unics[vet[i-1]] = true
		if vet[i - 1] == vet[i] && i < len(vet) - 1{
			count++
		} else {
			storageCount = append(storageCount, count)
			count = 1
		}
	}
	fmt.Println(storageCount)
	for u := range unics{
		teams = append(teams, Pair{One: u})
	}
	for _, v := range storageCount{
		teams = append(teams, Pair{Two: v})
	}

	
	
	return quickSort(teams, 0, len(freq) - 1)
}

func mnext(vet []int) []int {
	_ = vet
	return nil
}

func alone(vet []int) []int {
	_ = vet
	return nil
}

func couple(vet []int) int {
	_ = vet
	return 0
}

func hasSubseq(vet []int, seq []int, pos int) bool {
	_ = vet
	_ = seq
	_ = pos
	return false
}

func subseq(vet []int, seq []int) int {
	_ = vet
	_ = seq
	return -1
}

func erase(vet []int, posList []int) []int {
	_ = vet
	_ = posList
	return nil
}

func clear(vet []int, value int) []int {
	_ = vet
	_ = value
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "occurr":
			printSlice(occurr(str2vet(args[1])))
		case "teams":
			printSlice(teams(str2vet(args[1])))
		case "mnext":
			printSlice(mnext(str2vet(args[1])))
		case "alone":
			printSlice(alone(str2vet(args[1])))
		case "erase":
			printSlice(erase(str2vet(args[1]), str2vet(args[2])))
		case "clear":
			val, _ := strconv.Atoi(args[2])
			printSlice(clear(str2vet(args[1]), val))
		case "subseq":
			fmt.Println(subseq(str2vet(args[1]), str2vet(args[2])))
		case "couple":
			fmt.Println(couple(str2vet(args[1])))
		case "end":
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}

// Funções auxiliares

func str2vet(str string) []int {
	if str == "[]" {
		return nil
	}
	str = str[1 : len(str)-1]
	parts := strings.Split(str, ",")
	var vet []int
	for _, part := range parts {
		num, _ := strconv.Atoi(strings.TrimSpace(part))
		vet = append(vet, num)
	}
	return vet
}

func printSlice[T any](vet []T) {
	fmt.Print("[")
	for i, x := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(x)
	}
	fmt.Println("]")
}

func (p Pair) String() string {
	return fmt.Sprintf("(%v, %v)", p.One, p.Two)
}
