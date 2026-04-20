package main
import "fmt"





func main() {
    var n int
    fmt.Scan(&n)

    resto(n)
}

func resto(n int){
    if n <= 0{
        return
    }
    rest := n % 2
    n /= 2
    resto(n)
    fmt.Print(n, " ", rest, "\n")
}
