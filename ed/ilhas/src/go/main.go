package main

import (
	"bufio"
	"fmt"
	"os"
)

// Não modifique a assinatura da função numIslands
// Ela é a função que será chamada no LeetCode para resolver o problema
func numIslands(grid [][]byte) int {
    numIsle := 0
    
	var isle func(grid [][]byte, r, c int) bool
	isle = func(grid [][]byte, r, c int) bool {

		destruct := false
		if (r < 0 || r >= len(grid)) || (c < 0 || c >= len(grid[0])) {
			return true
		}
		if destruct {
			return true
		}

		if grid[r][c] == '1' {
            grid[r][c] = '0'
			destruct = true
			isle(grid, r+1, c)
			isle(grid, r-1, c)
			isle(grid, r, c+1)
			isle(grid, r, c-1)
			return true
		}
		return false
	}
	rows := len(grid)
	cowls := len(grid[0])
    for r := range rows{
        for c := range cowls{
            if grid[r][c] == '1'{
                numIsle++
                isle(grid, r, c)
            }
        }
    }

	return numIsle

}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	grid := make([][]byte, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		grid[i] = []byte(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	result := numIslands(grid)
	fmt.Println(result)
}
