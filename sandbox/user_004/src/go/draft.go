package main

import (
	"fmt"
	"strings"
)


type Info struct {
    QtdVoga int
    QtdCons int
    QtdPala int
}



func main() {
    texto := "  eu me  lasquei"

    info := Analisar(texto)

    fmt.Printf("Vogais: %d\n", info.QtdVoga)
    fmt.Printf("Consoantes: %d\n", info.QtdCons)
    fmt.Printf("Palavras: %d\n", info.QtdPala)
}


func Analisar(texto string) Info {
    var info Info
    runes := []rune(texto)

    for _, c := range runes {
        if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'{
            info.QtdVoga++
        } else if (c != ' '){
            info.QtdCons++
        }
    }

    palavras := strings.Split(texto, " ")
    info.QtdPala = len(palavras)

    return info
}