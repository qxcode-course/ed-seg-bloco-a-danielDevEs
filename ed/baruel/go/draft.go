package main
import "fmt"
func main() {
    var album, possui int
    fmt.Scanf("%d\n%d", &album, &possui)

    figurinhas := make([]int, possui)

    for i := 0; i < possui; i++ {
        fmt.Scan(&figurinhas[i])
    }

    repetidas := []int{}
    flag := false
    for i := 0; i < len(figurinhas); i++ {
        flag = false
        for j := 1; j < len(figurinhas) - i; j++ {
            if figurinhas[i] == figurinhas[i + j] {
                if flag == true {
                    break
                }
                repetidas = append(repetidas, figurinhas[i])
                flag = true
            }
        }
    }

    if len(repetidas) == 0 {
        fmt.Println("N")
    } else {
        for i, r := range repetidas {
            fmt.Print(r)
            if i < len(repetidas) - 1 {  
                fmt.Print(" ")
            }
        }
        fmt.Print("\n")
    }

    faltando := []int{}
    flag2 := false
    for i := 1; i <= album; i++ {
        flag2 = false
        for j := 0; j < len(figurinhas); j++ {
            if i == figurinhas[j] {
                flag2 = true
            }
        }
        if flag2 == false {
            faltando = append(faltando, i)
        }
    }

    if len(faltando) == 0 {
        fmt.Println("N")
    } else {
        for i, f := range faltando {
            fmt.Print(f)
            if i < len(faltando) - 1 {  
                fmt.Print(" ")
            }
        }
        fmt.Print("\n")
    }
}
