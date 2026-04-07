package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func soma(vet []int) []int{
	soma := []int{}
	for v := 0; v < len(vet) - 1; v++ {
		soma = append(soma, vet[v] + vet[v + 1])
	}
	return soma
}
func processa(vet []int) {
	aux := []int{}
	if len(vet) > 0 {
		aux = soma(vet)
		processa(aux)
		fmt.Println(Join(vet, " "))
		vet = aux
	}

	// 1. defina o ponto de parada
	// 2. monte o vetor auxiliar com os resultados das somas
	// 3. chame recursivamente a função processa para o vetor auxiliar
	// 4. imprima o vetor original
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	line := scanner.Text()
	parts := strings.Fields(line)
	vet := []int{}
	for _, part := range parts {
		if value, err := strconv.Atoi(part); err == nil {
			vet = append(vet, value)
		}
	}
	processa(vet)
}

func Join[T any](v []T, sep string) string {
	if len(v) == 0 {
		return ""
	}
	s := "[ "

	for i, x := range v {
		if i > 0 {
			s += sep
		}
		s += fmt.Sprintf("%v", x)
	}
	s += " ]"
	return s
}
