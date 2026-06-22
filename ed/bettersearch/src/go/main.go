package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)



func BetterSearch(slice []int, value int) (bool, int) {
	if len(slice) == 0{
		return false, 0
	}
	mid := 0
	low := 0
	high := len(slice) - 1

	for low <= high {
		mid = low + (high-low) / 2

		if slice[mid] == value {
			return true, mid
		}

		if slice[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if mid < value {
		return false, mid + 1
	}
	return false, mid
} 

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	slice := []int{}
	for _, elem := range parts[1 : len(parts)-1] {
		value, _ := strconv.Atoi(elem)
		slice = append(slice, value)
	}
	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	found, result := BetterSearch(slice, value)
	if found {
		fmt.Println("V", result)
	} else {
		fmt.Println("F", result)
	}
}
