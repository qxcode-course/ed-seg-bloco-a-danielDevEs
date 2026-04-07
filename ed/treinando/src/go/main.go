package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func tostr(vet []int) string {
	str := ""
	i := 0
	if i == len(vet) - 1 || len(vet) == 0{
		return "[" + str + "]\n"
	}
	str += fmt.Sprint(vet[i])
	str += ","
	i++
	tostr(vet)
	return str
}

func tostrrev(vet []int) string {
	_ = vet
	return ""
}

// reverse: inverte os elementos do slice
func reverse(vet []int) []int {
	if len(vet) == 1 || len(vet) == 0{
		return vet
	}
	for v := range vet{
		vet[v], vet[len(vet) - v] = vet[len(vet) - v], vet[v]
	}
	reverse(vet)
	return vet
}

// sum: soma dos elementos do slice
func sum(vet []int) int {
	_ = vet
	return 0
}

// mult: produto dos elementos do slice
func mult(vet []int) int {
	_ = vet
	return 0
}

// min: retorna o índice e valor do menor valor
// crie uma função recursiva interna do modelo
// var rec func(v []int) (int, int)
// para fazer uma recursão que retorna valor e índice
func min(vet []int) int {
	_ = vet
	return 0
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
