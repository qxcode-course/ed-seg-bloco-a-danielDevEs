package main

import (
	"fmt"
)



func apodrecerLaranjas(grid [][]int) int{
    maxR := len(grid)
    maxC := len(grid[0])
    
    var apodrecer func(r, c int) bool
    // segundos := 0

    
    apodrecer = func(r, c int) bool {
        
        if (r < 0 || r >= maxR) || (c < 0 || c >= maxC){
            return false
        }
        
        
        
        laranja := grid[r][c]
        
        if laranja == 1 {
            grid[r][c] = 9
        }

        dirs := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
        
        for _, d := range dirs{
            nr, nc := r+d[0], c+d[1]
            apodrecer(nr, nc)
        }
        return true
    }

    for r := 0; r < maxR; r++ {
        for c := 0; c < maxC; c++ {
            if grid[r][c] == 2 {
                apodrecer(r, c)
            }
        }
    }
    
    for r := 0; r < maxR; r++ {
        for c := 0; c < maxC; c++ {
            if grid[r][c] == 9 {
                grid[r][c] = 2
            }
        }
    }
return 0
}






func main() {
    var maxR, maxC int
    fmt.Scanf("%d%d", &maxR, &maxC)

    grid := make([][]int, maxR)

    for r := 0; r < maxR; r++ {
        grid[r] = make([]int, maxC)
    }  


    for r := 0; r < maxR; r++ {
        for c := 0; c < maxC; c++ {
            fmt.Scan(&grid[r][c])
        }
    }
    apodrecerLaranjas(grid)
    fmt.Println(grid)
}