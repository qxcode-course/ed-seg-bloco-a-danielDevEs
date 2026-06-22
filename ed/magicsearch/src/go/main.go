package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MagicSearch(slice []int, value int) int {
	low := 0
	mid := 0
	high := len(slice) - 1
	if len(slice) == 0 {
		return 0
	}

	for low <= high {
		mid = low + (high - low)/2

		if value == slice[mid] {
			for {
				if mid == len(slice)-1{
					return mid
				}
				if slice[mid+1] == value {
					mid++
				} else {
					break
				}
			}
			return mid
		}

		if slice[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if slice[mid] < value{
		return mid + 1
	} 
	return mid
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	slice := make([]int, 0, 1)
	for _, elem := range parts[1 : len(parts)-1] {
		value, _ := strconv.Atoi(elem)
		slice = append(slice, value)
	}

	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	result := MagicSearch(slice, value)
	fmt.Println(result)
}
