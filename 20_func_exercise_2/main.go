package main

import (
	"fmt"
)

func main() {
	s := []int{
		10,
		20,
		30,
		50,
		100,
		250,
		1000,
		7500,
	}
	x := foo(s...)
	y := bar(s)

	fmt.Println(x, y)

}

func foo(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func bar(x []int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
