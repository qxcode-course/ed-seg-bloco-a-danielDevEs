 package main

import (
	"bufio"
	"fmt"
	"os"
)

type Position struct {
	row, cowl int
}


func PodeCercar(board [][]byte) map[Position]bool {
	visitados := make(map[Position]bool)
	for i := range len(board) {
			if board[0][i] == 'O'{
				visitados[Position{row: 0, cowl: i}] = true
			}
		}
		for i := range len(board[0]) {
			if board[i][0] == 'O'{
				visitados[Position{row: i, cowl: 0}] = true
			}
		}
		for i := range len(board) {
			if board[len(board)-1][i] == 'O'{
				visitados[Position{row: len(board)-1, cowl: i}] = true
			}
		}
		for i := range len(board[0]) {
			if board[i][len(board[0])-1] == 'O'{
				visitados[Position{row: i, cowl: len(board[0])-1}] = true
			}
		}
	return visitados
}

// NÃO ALTERE A ASSINATURA DA FUNÇÃO solve
func solve(board [][]byte) {

	visitados := make(map[Position]bool)
	visitados = PodeCercar(board)
	

	var cercar func (board [][]byte, r, c int) bool

	cercar = func(board [][]byte, r, c int) bool {
		if (r < 0 || r >= len(board)) || (c < 0 || c >= len(board[0])) || board[r][c] == 'X'{
			return false
		}

		if board[r][c] == 'O'{
			board[r][c] = '*'
			cercar(board, r+1, c)
			cercar(board, r-1, c)
			cercar(board, r, c+1)
			cercar(board, r, c-1)
			return true
		}
		return false
	}
	var maxR, maxC int = len(board), len(board[0])
	for r := range maxR {
		for c := range maxC {
			if board[r][c] == 'X'{
				continue
			}
			if visitados[Position{row: r, cowl: c}]{
				cercar(board, r, c)
			}

		}
	}

	for r := range maxR {
		for c := range maxC {
			if board[r][c] == 'X'{
				continue
			}
			if board[r][c] == 'O'{
				board[r][c] = 'X'
			}
			if board[r][c] == '*'{
				board[r][c] = 'O'
			}
		}
	}
}
	


// NÃO ALTERE A MAIN
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var nrows, ncols int
	fmt.Sscanf(scanner.Text(), "%d %d", &nrows, &ncols)
	board := make([][]byte, nrows)
	for i := 0; i < nrows; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}
	solve(board)
	for _, row := range board {
		fmt.Println(string(row))
	}
}
