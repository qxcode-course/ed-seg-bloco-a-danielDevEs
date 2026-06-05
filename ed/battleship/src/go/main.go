package main

import (
	"bufio"
	"fmt"
	"os"
)

// Função que será chamada no LeetCode
func countBattleships(board [][]byte) int {
	//

	var qtdBattleShips func(r, c int) bool

	qtdBattleShips = func(r, c int) bool {
		
		if (r < 0 || r >= len(board)) || (c < 0 || c >= len(board[0])) || board[r][c] != 'X' {
			return false
		}

		board[r][c] = '*'

		dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

		for _, d := range dirs{
			nr, nc := r+d[0], c+d[1]
			qtdBattleShips(nr, nc)
		}
		return true
	}
	maxR, maxC := len(board), len(board[0])
	ships := 0
	for r := range maxR {
		for c := range maxC {
			if board[r][c] == 'X'{
				ships++
				qtdBattleShips(r, c)
			}
		}
	}
	return ships
	
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)

	board := make([][]byte, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}

	result := countBattleships(board)
	fmt.Println(result)
}
