package main
import "fmt"
func main() {
    var size, pos int
    fmt.Scan(&size, &pos)
    
    elementos := make([]int, size)
    for i := 0; i < len(elementos); i++{
        elementos[i] = i + 1
    }
    
    pos--
    
    vivos := size
    for {
        toString(elementos, pos)
        
        if(vivos == 1){
         break
        }

        morto := procurar_vivo(elementos, pos)
        elementos[morto] = 0

        vivos--
       
        pos = procurar_vivo(elementos, pos)
    }
}

func procurar_vivo(elementos []int, pos int) int {
    for {
        pos = (pos + 1) % len(elementos)
        
        if elementos[pos] != 0 {
            return pos
        }
    }
}

func toString(elementos []int, pos int) {
    fmt.Print("[ ")
    for i := 0; i < len(elementos); i++ {
        if elementos[i] != 0 {
            if i == pos {
                fmt.Printf("%d> ", elementos[i])
            } else {
                fmt.Printf("%d ", elementos[i])
            }
        }
    }
    fmt.Print("]\n")
}