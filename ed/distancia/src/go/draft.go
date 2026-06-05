package main

import (
	"fmt"
)



func verifySpace(instance []byte, limit, index int, num byte) bool{
    space := true

        countL := 0
        for i := index-1; i >= 0; i--{
            if countL == limit{
                break
            }
            if instance[i] == num{
                space = false
            }
            countL++
        }
        countR := 0
        for i := index+1; i < len(instance); i++{
            if countR == limit{
                break
            }
            if instance[i] == num{
                space = false
            }
            countR++
        }

        return space
}








func main() {
    var instance string
    var limit int
    
    fmt.Scanln(&instance)
    fmt.Scanln(&limit)


    numbers := []byte{}

    for i := 0; i <= limit; i++{
        
        numbers = append(numbers, byte('0'+i))
    }

    pontos := 0
    for _, e := range instance{
        if e == '.'{
            pontos++
        }
    }
    // validsBlocks := [][pontos]valids
    
    copyInstance := []byte(instance)
    
    
    
    
    var validCombinations func(copyInstance []byte, index int) (bool, []byte)
    
    validCombinations = func(copyInstance []byte, index int) (bool, []byte) {

        if index == len(copyInstance) {
            return true, copyInstance
        }

        if copyInstance[index] != '.' {
            return validCombinations(copyInstance, index+1)
        }

        for _, num := range numbers {

            if !verifySpace(copyInstance, limit, index, num) {
                continue
            }

            // escolhe
            copyInstance[index] = num

            // explora
            ok, result := validCombinations(copyInstance, index+1)

            if ok {
                return true, result
            }

            // desfaz
            copyInstance[index] = '.'
        }

        return false, copyInstance
    }
    
    flag := false

        for i, e := range instance{
           if e == '.'{
               status, newInstance := validCombinations(copyInstance, i)
               if status{
                    instance = string(newInstance)
                    flag = true
                    break
               }
           }
        }
        if flag {
            fmt.Println(instance)
        }
}