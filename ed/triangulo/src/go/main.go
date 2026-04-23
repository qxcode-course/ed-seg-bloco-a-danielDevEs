package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)



func processa(vet []int) {
	if(len(vet) == 1){
		fmt.Println(Join(vet, " "))
		return
	}
	aux := []int{}
	for i := 0; i < len(vet)-1; i++ {
		aux = append(aux, vet[i] + vet[i+1])
	}
	processa(aux)
	fmt.Println(Join(vet, " "))
	

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
