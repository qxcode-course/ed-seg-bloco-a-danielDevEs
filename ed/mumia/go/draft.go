package main

import "fmt"

func main() {
	var nome string
	var idade int
	fmt.Scan(&nome, &idade)

	var tipo string
	if idade < 12 {
		tipo = "crianca"
	} else if idade < 18 {
		tipo = "jovem"
	} else if idade < 65 {
		tipo = "adulto"
	} else if idade < 1000 {
		tipo = "idoso"
	} else {
		tipo = "mumia"
	}

	fmt.Println(nome + " eh " + tipo)

}
