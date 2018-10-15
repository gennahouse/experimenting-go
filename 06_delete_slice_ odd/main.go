package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Slice before delete odds values:", slice)
	val := 0
	for i, v := range slice {
		if v%2 != 0 {
			val = i + 1
			slice = append(slice[:i], slice[val:]...)
		}
	}
	fmt.Println("Slice after deliting odds number:", slice)
}
