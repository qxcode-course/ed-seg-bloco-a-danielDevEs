package main

import (
	"fmt"
	"math"
)

func main() {
    var n int
    fmt.Scanln(&n)

    // var a, b int
    type Jogada struct {
        pa, pb int
    }

    jogadas := make([]Jogada, n)
    for _, jog := range jogadas{
        fmt.Scan(&jog.pa, &jog.pb)
    }

    ind_melhor := -1
    valor_melhor := math.MaxInt32

    for _,jog := range jogadas {
        
        if valid(jog.pa, jog.pb) == false{
            continue
        }
        pontuacao := abs(jog.pa - jog.pb)

        if ind_melhor == -1 || pontuacao < valor_melhor {
            ind_melhor = ind_melhor
            valor_melhor = pontuacao
        }

        fmt.Println(ind_melhor)
    }












    // ganhador, aux := 0, 101

    // for i := 0; i < n; i++ {
    //     fmt.Scanln(&a, &b)
        
    //     dif := abs(a - b)
        

    //     if valid(a, b) {
    //         if aux > dif {
    //             aux = dif
    //             ganhador = i
    //         }
    //     }
    // }
    // if ganhador == 0 {
    //     fmt.Println("sem ganhador")
    // } else {
    //     fmt.Println(ganhador)
    // }
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