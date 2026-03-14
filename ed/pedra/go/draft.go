package main

import (
	"fmt"
)

func main() {
    var n int
    fmt.Scanln(&n)

    var a, b int
    ganhador, aux := 0, 101

    for i := 0; i < n; i++ {
        fmt.Scanln(&a, &b)
        
        dif := abs(a - b)
        

        if valid(a, b) {
            if aux > dif {
                aux = dif
                ganhador = i
            }
        }
    }
    if ganhador == 0 {
        fmt.Println("sem ganhador")
    } else {
        fmt.Println(ganhador)
    }
}

func valid(a, b int) bool {
    return (a >= 10 && a <= 100) &&
           (b >= 10 && b <= 100)
}

func abs(dif int) int {
    if dif < 0 {
        return -dif
    }
    return dif
}