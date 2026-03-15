package main
import "fmt"
func main() {
    var qtdgomos int
    var direction string

    fmt.Scanln(&qtdgomos, &direction)

    // cobra := make([][]int, qtdgomos)
    // for i := range cobra{
    //     cobra[i] = make([]int, qtdgomos)
    // }
    
    lineX, lineY := make([]int, qtdgomos), make([]int, qtdgomos)
    // var posX, posY int

    for i := range qtdgomos{
        fmt.Scan(&lineX[i], &lineY[i])
    }
    // fmt.Println(lineX)
    // fmt.Println(lineY)

    for i := range len(lineX) - 1{
        lineX[qtdgomos - i -1] = lineX[qtdgomos - i - 2]
    }
    for i := range len(lineY) - 1{
                lineY[qtdgomos - i - 1] = lineY[qtdgomos - i - 2]
    }

    switch direction {
        case "R":
            lineX[0] += 1
        case "L":
            lineX[0] -= 1
        case "U":
            lineY[0] -= 1
        case "D":
            lineY[0] += 1
    }

    for i := range qtdgomos {
        fmt.Printf("%d %d\n", lineX[i], lineY[i])
    }
    

}
