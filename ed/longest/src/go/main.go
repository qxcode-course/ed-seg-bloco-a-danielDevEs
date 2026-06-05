package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	linhas := len(matrix)
	colunas := len(matrix[0])

	// Matriz de memorização (cache) para guardar os resultados já calculados
	memo := make([][]int, linhas)
	for i := range memo {
		memo[i] = make([]int, colunas)
	}

	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	// DFS que RETORNA um número (o tamanho do caminho) a cada avanço
	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		// Se já calculamos o caminho a partir desta célula antes, retorna o cache
		if memo[r][c] != 0 {
			return memo[r][c]
		}

		maxCaminho := 1 // O caminho mínimo para qualquer célula é 1 (ela mesma)

		// Explora as 4 direções
		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]

			// Verifica os limites e se o vizinho é estritamente maior
			if nr >= 0 && nr < linhas && nc >= 0 && nc < colunas && matrix[nr][nc] > matrix[r][c] {
				// Aqui está o "retorno a cada avanço": calculamos o tamanho vindo do vizinho
				caminhoDoVizinho := 1 + dfs(nr, nc)
				
				// Queremos o MAIOR caminho possível, então comparamos as opções
				if caminhoDoVizinho > maxCaminho {
					maxCaminho = caminhoDoVizinho
				}
			}
		}

		// Guarda o maior resultado encontrado para esta célula no cache
		memo[r][c] = maxCaminho
		return maxCaminho
	}

	maiorCaminhoGlobal := 0

	// Como o caminho pode começar em qualquer ponto da matriz, testamos todos
	for r := 0; r < linhas; r++ {
		for c := 0; c < colunas; c++ {
			caminhoAtual := dfs(r, c)
			if caminhoAtual > maiorCaminhoGlobal {
				maiorCaminhoGlobal = caminhoAtual
			}
		}
	}

	return maiorCaminhoGlobal
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}
	parts := strings.Fields(scanner.Text())
	if len(parts) < 2 {
		return
	}
	nl, _ := strconv.Atoi(parts[0])
	nc, _ := strconv.Atoi(parts[1])

	matrix := make([][]int, nl)
	for i := 0; i < nl; i++ {
		if !scanner.Scan() {
			return
		}
		tokens := strings.Fields(scanner.Text())
		row := make([]int, nc)
		for j := 0; j < nc && j < len(tokens); j++ {
			v, _ := strconv.Atoi(tokens[j])
			row[j] = v
		}
		matrix[i] = row
	}

	fmt.Println(longestIncreasingPath(matrix))
}
