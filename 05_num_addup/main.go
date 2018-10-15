package main

import (
	"fmt"
)

func main() {
	result := sumNums(10)
	fmt.Println(result)
}

func sumNums(i int) int {
	res := 0
	for j := 0; j <= i; j++ {
		res += j
	}
	return res
}
