package main
import "fmt"
func main() {
    qtd := 0
    fmt.Scan(&qtd)
    fmt.Println(gerarFib(qtd))
    
}

func gerarFib(qtdElem int) []int {
    var fib []int = make([]int, qtdElem)

    for i := 0; i < qtdElem; i++{
        if i < 2 {
            fib[i] = 1
            continue
        }
        fib[i] = fib[i-1] + fib[i-2]
    }

    return fib
}