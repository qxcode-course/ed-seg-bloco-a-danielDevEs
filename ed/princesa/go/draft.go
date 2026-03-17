package main
import "fmt"
func main() {
    var size, pos int
    fmt.Scan(&size, &pos)

    vivos := make([]int, size)
    for i := 1; i < len(vivos); i++{
        vivos[i - 1] = i
    }

    fmt.Println(vivos)
    i := 0
    for {
        if vivos[i] == 0 {
            continue
        }
        if vivos[i] == pos{
            prox := vivos[i + 1] % size
            vivos[prox] = 0
        }
        i++
        if i > size {
            i = i % size
        }
        fmt.Println(vivos)
    }
}
    