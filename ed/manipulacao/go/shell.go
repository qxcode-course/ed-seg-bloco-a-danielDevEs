package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getMen(vet []int) []int {
	homens := []int{}
	for i := range vet{
		if vet[i] > 0{
			homens = append(homens, vet[i])
		}
	}
	return homens
}

func getCalmWomen(vet []int) []int {
	mulheres_calmas := []int{}
	for v := range vet{
		if vet[v] < 0 && vet[v] > -10{
			mulheres_calmas = append(mulheres_calmas, vet[v])
		}
	}
	return mulheres_calmas
}

func sortVet(vet []int) []int {
	return quickSortOriginal(vet, 0, len(vet) - 1)
}

func sortStress(vet []int) []int {
	return quickSortAlterado(vet, 0, len(vet) - 1)
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func reverse(vet []int) []int {
	invert := []int{}
	for i := len(vet) - 1; i >= 0; i--{
		invert = append(invert, vet[i])
	}
	return invert
}

func quickSortAlterado(v []int, low, high int) []int{
	if low < high {
		p := partitionAlterada(v, low, high)

		quickSortAlterado(v, low, p-1)
		quickSortAlterado(v, p+1, high)
	}
	return v
}

func quickSortOriginal(v []int, low, high int) []int{
	if low < high {
		p := partitionOriginal(v, low, high)

		quickSortOriginal(v, low, p-1)

		quickSortOriginal(v, p+1, high)

	}
	return v
}
func partitionOriginal(v []int, low, high int) int{
	pivot := v[high]
	i := low - 1

	for j := low; j < high; j++ {
		if v[j] < pivot {
			i++
			v[i], v[j] = v[j], v[i]
		}
	}

	v[i+1], v[high] = v[high], v[i+1]
	return i+1
}

func partitionAlterada(v []int, low, high int) int{
	pivot := v[high]
	i := low - 1

	for j := low; j < high; j++ {
		if abs(v[j]) < abs(pivot) {
			i++
			v[i], v[j] = v[j], v[i]
		}
	}

	v[i+1], v[high] = v[high], v[i+1]

	return i+1
}

func unique(vet []int) []int {
	freq := make(map[int]int)
	unicos := []int{}
   
	for _, num := range vet{
		if freq[num] == 0 {
			unicos = append(unicos, num)
		}
		freq[num]++
	}

	return unicos
}

func repeated(vet []int) []int {
	freq := make(map[int]int)
	repetidos := []int{}
	for _, num := range vet{
		if freq[num] >= 1 {
			repetidos = append(repetidos, num)
		}
		freq[num]++
	}
	return repetidos
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}

