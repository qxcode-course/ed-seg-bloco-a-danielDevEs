package main

import (
	"fmt"
	"slices"
)



func main() {

    var n, target int
    fmt.Scan(&n, &target)

    var numbers = make([]int, n)
    for i := range numbers{
        fmt.Scan(&numbers[i])
    }
    slices.Sort(numbers)

    var canFindSum func(index int, currentSum int) bool

    canFindSum = func(index int, currentSum int) bool {

        if currentSum == target {
            return true
        }
        if currentSum > target {
            return false
        }
        if index >= len(numbers) {
            return false 
        }

        
        // Escolha A: Incluir o número atual na soma
        include := canFindSum(index + 1, currentSum + numbers[index])
        if include {
            return true // Se esse caminho deu certo, não precisa testar mais nada!
        }

        // Escolha B: Ignorar o número atual
        exclude := canFindSum(index + 1, currentSum)
        
        return exclude
    }
    fmt.Println(canFindSum(0, 0))
}
