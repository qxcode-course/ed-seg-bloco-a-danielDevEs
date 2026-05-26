package main

import (
	"fmt"
)


func main() {
    
    var instance string
    fmt.Scan(&instance)
    // VetString := []byte(instance)
    // for {
    //     var element rune
    //     fmt.Scan(&element)
    //     if element == '\n'{
    //         break
    //     }
    //     instance = append(instance, element)
    // }
    // fmt.Println(string(VetString))
    var instancecopy string
    instancecopy = instance

    var limit int
    fmt.Scan(&limit)
    numbers := make([]int, limit+1)

    for i := 0; i <= limit; i++{
        numbers[i] = i
    }


    var PutNumber func(instance []rune, numbers []int, idx int) []rune

    PutNumber = func(instance []rune, numbers []int, idx int) []rune {

        if idx == len(instance){
            // fmt.Println("parou")
            // for _, v := range instance{
            //     if v == '.'
            //     return PutNumber(in)
            // }
            return instance
        }
        
        if instance[idx] != '.'{
            // fmt.Println("NUMERO")
            return PutNumber(instance, numbers, idx+1)
        }
        // fmt.Println("Ponto")

        colocado := false

        for _, n := range numbers{
            if colocado == true {
                break
            }

            instance[idx] = rune(n+48)
            colocado = true
            // fmt.Println(string(instance))

            var controlLeft int = 0
            for instanceIndex := idx-1; instanceIndex >= 0; instanceIndex--{
                if controlLeft == limit{
                    break
                }
                if instance[instanceIndex] == '.'{
                    controlLeft++
                    continue
                }
                // fmt.Println("TRAS: index:", string(instance[instanceIndex]), "idx:", string(instance[idx]))
                if instance[instanceIndex] == instance[idx]{
                    instance[idx] = '.'
                    // fmt.Println("tirado")
                    colocado = false
                    break
                }
                controlLeft++
            }
    
    
            var controlRight int =  0
            for instanceIndex := idx+1; instanceIndex < len(instance); instanceIndex++{
                if controlRight == limit{
                    break
                }
                if instance[instanceIndex] == '.'{
                    controlRight++
                    continue
                }
                // fmt.Println("FRENTE: index:", string(instance[instanceIndex]), "idx:", string(instance[idx]))
                if instance[instanceIndex] == instance[idx]{
                    instance[idx] = '.'
                    // fmt.Println("tirado")
                    colocado = false
                    break
                }
                controlRight++
            }
        }
        return PutNumber(instance, numbers, idx+1)
    }

    fmt.Println(string(PutNumber([]rune(instance), numbers, 0)))

}
