package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func str(vet []int, i int) string{
	stri := ""
	if i == len(vet){
		return stri
	}
	stri += fmt.Sprint(vet[i])
	if(i < len(vet) - 1){
		stri += ", "
	}
	return stri + str(vet, i + 1)

}

func tostr(vet []int) string {
	return "[" + str(vet, 0) + "]"
}

func strIn(vet []int, i int) string{
	stri := ""
	if i == -1{
		return stri
	}
	stri += fmt.Sprint(vet[i])
	if(i > 0){
		stri += ", "
	}
	return stri + strIn(vet, i - 1)

}
func tostrrev(vet []int) string {
	return "[" + strIn(vet, len(vet)- 1) + "]"
}

// reverse: inverte os elementos do slice
func reversevet(vet []int, i, f int) []int{
	if i > f{
		return vet
	}
	vet[i], vet[f] = vet[f], vet[i]
	return reversevet(vet, i+1, f-1)
}
func reverse(vet []int) []int {
	return reversevet(vet, 0, len(vet) - 1)
}

// sum: soma dos elementos do slice

func sum(vet []int) int {
	if len(vet) == 0{
		return 0
	}
	if len(vet) == 1{
		return vet[0]
	}
	return vet[0] + sum(vet[1:])
}

// mult: produto dos elementos do slice
func mult(vet []int) int {
	if len(vet) == 0{
		return 1
	}
	if len(vet) == 1{
		return vet[0]
	}
	return vet[0] * mult(vet[1:])
}

// min: retorna o índice e valor do menor valor
// crie uma função recursiva interna do modelo
// var rec func(v []int) (int, int)
// para fazer uma recursão que retorna valor e índice
func min(vet []int) int {
	if len(vet) == 0 {
		return -1
	}

	var rec func(i int) (int, int)

	rec = func(i int) (int, int) {
		if i == len(vet)-1 {
			return i, vet[i]
		}

		idx, val := rec(i + 1)

		if vet[i] <= val {
			return i, vet[i]
		}

		return idx, val
	}

	idx, _ := rec(0)
	return idx
}

func main() {
	var vet []int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			return
		case "read":
			vet = nil
			for _, arg := range args[1:] {
				if val, err := strconv.Atoi(arg); err == nil {
					vet = append(vet, val)
				}
			}
		case "tostr":
			fmt.Println(tostr(vet))
		case "torev":
			fmt.Println(tostrrev(vet))
		case "reverse":
			reverse(vet)
		case "sum":
			fmt.Println(sum(vet))
		case "mult":
			fmt.Println(mult(vet))
		case "min":
			fmt.Println(min(vet))
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
