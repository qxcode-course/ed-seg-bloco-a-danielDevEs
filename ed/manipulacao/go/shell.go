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
	aux := 0
	for i := 0; i < len(vet) - 1 -i; i++{
		if vet[i] < vet[i+1]{
			aux = vet[i]
			vet[i] = vet[i+ 1]
			vet[i + 1] = aux
		}
	}
	return vet
}

func sortStress(vet []int) []int {
	_ = vet
	return vet
}

func reverse(vet []int) []int {
	invert := []int{}
	for i := len(vet) - 1; i >= 0; i--{
		invert = append(invert, vet[i])
	}
	return invert
}

func unique(vet []int) []int {
	unicos := []int{}
	unicos = append(unicos, vet[0])
	for i := 1; i < len(vet) - 1; i++{
		for j := 0; j < len(unicos); j++{
			if(vet[i] != vet[j]){
				unicos = append(unicos, vet[i])
			}
		}
	}
	return unicos
}

func repeated(vet []int) []int {
	_ = vet
	return vet
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

