package main
import "fmt"
func main() {
    var n int
    fmt.Scanln(&n)

    qtd := make([]int, n)
    descasdos := []int{}

    for i := 0; i < len(qtd); i++ {
        fmt.Scan(&qtd[i])
        
        if(qtd[i] > 0){
            descasdos = append(descasdos, qtd[i])
        }
    }
    
    flag := false
    count := 0

    for i := 0; i < len(qtd); i++{
        flag = false
        for j := 0; j < len(descasdos); j++{
            if flag == true{
                break
            }

            if qtd[i] < 0 {
                if (descasdos[j] == (qtd[i] * -1)) {
                    descasdos[j] = 0
                    flag = true
                    count++
                }
            }
        }
    }

    fmt.Println(count)
}
