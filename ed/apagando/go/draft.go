package main
import "fmt"
func main() {
	var qtd_pessoas int
	fmt.Scan(&qtd_pessoas)

	fila := make([]int, qtd_pessoas)
	identi := 0
	for i := 0; i < qtd_pessoas; i++{
		fmt.Scan(&identi)
		fila[i] = identi
	}

	var qtd_deixaram int
	fmt.Scan(&qtd_deixaram)

	fila_deixada := make([]int, qtd_deixaram)
	identi_Deixado := 0
	for i := 0; i < qtd_deixaram; i++{
		fmt.Scan(&identi_Deixado)
		fila_deixada[i] = identi_Deixado
	}

	fmt.Println(ToString(retirar(fila, fila_deixada)))
  
}

func retirar(fila, deixaram []int) []int {
	valores := make(map[int]bool)
	for _, v := range deixaram {
		valores[v] = true
	}

	fila_atualizada := []int{}
	for _, v := range fila {
		if !valores[v] {
			fila_atualizada = append(fila_atualizada, v)
		}
	}

	return fila_atualizada
}

func ToString(fila_atualizada []int) string{
	str := ""
	for i := 0; i < len(fila_atualizada); i++ {
		str += fmt.Sprint(fila_atualizada[i])
		str += " "
	}
	
	return str
}

