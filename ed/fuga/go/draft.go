package main

import "fmt"

func main() {
	var helico, polic, fugit, direc int
	fmt.Scan(&helico, &polic, &fugit, &direc)

	for {

		fugit = (fugit + direc + 16) % 16

		if fugit == polic {
			fmt.Println("N")
			break
		}
		if fugit == helico {
			fmt.Println("S")
			break
		}
	}
}
