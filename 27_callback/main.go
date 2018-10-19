package main

import (
	"fmt"
)

var sum int

func main() {
	x := callB(count, 13, 14)
	fmt.Println(x)
}

func count(i, n int) int {
	sum = i + n
	return sum
}

func callB(f func(i, n int) int, z int, y int) int {
	x := f(z, y)
	return x
}
