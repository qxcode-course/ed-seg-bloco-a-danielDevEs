package main

import (
	"fmt"
	// "os"
	// "bufio"
)
    
// func Éreservada(char rune) bool {
//     reservadas := []rune{'\n', '<', '>', 'D', 'B', 'R'}

//     for _, c := range reservadas{
//         if char == c {
//             return true
//         }
//     }

//     return false
// }

func main() {
    frase := NewList()

    cursor := frase.root
    for {
        var char rune

        fmt.Scanf("%c", &char)

        if char == '\r' {
            continue
        }

        if char == '\n'{
            break
        }

		switch char {

        case 'R':
            frase.Insert(cursor, '\n')

		case '<':
			if cursor.prev != frase.root {
				cursor = cursor.prev
			}

		case '>':
			if cursor != frase.root {
				cursor = cursor.next
			}

		case 'B':
			if cursor.prev != frase.root {
				frase.Erase(cursor.prev)
			}

		case 'D':
			if cursor != frase.root {
				next := cursor.next
				frase.Erase(cursor)
				cursor = next
			}

		default:
			frase.Insert(cursor, char)
		}
	}

    for node := frase.Front(); node != frase.root; node = node.next {
		if node == cursor {
			fmt.Print("|")
		}
		fmt.Printf("%c", node.data)
	}
	if cursor == frase.root {
		fmt.Print("|")
	}
    fmt.Println()
}
