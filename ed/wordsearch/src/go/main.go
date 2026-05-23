package main

import (
	"bufio"
	"fmt"
	"os"
)

func exist(grid [][]byte, word string) bool {

	rows, cols := len(grid), len(grid[0])

	type Frame struct {
		r, c  int
		index int
		path  map[[2]int]bool
	}

	stack := NewStack[Frame]()

	for r := range rows{
		for c := range cols {

			if grid[r][c] == word[0] {

				path := make(map[[2]int]bool)
				path[[2]int{r, c}] = true

				stack.Push(Frame{
					r:     r,
					c:     c,
					index: 0,
					path:  path,
				})
			}
		}
	}

	directions := [][2]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	for !stack.IsEmpty() {

		current := stack.Pop()

		if current.index == len(word)-1 {
			return true
		}

		for _, d := range directions {

			nr := current.r + d[0]
			nc := current.c + d[1]

			if nr < 0 || nr >= rows ||
				nc < 0 || nc >= cols {
				continue
			}

			if current.path[[2]int{nr, nc}] {
				continue
			}

			if grid[nr][nc] != word[current.index+1] {
				continue
			}

			newPath := make(map[[2]int]bool)

			for k, v := range current.path {
				newPath[k] = v
			}

			newPath[[2]int{nr, nc}] = true

			stack.Push(Frame{
				r:     nr,
				c:     nc,
				index: current.index + 1,
				path:  newPath,
			})
		}
	}

	return false
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]byte, 0)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
