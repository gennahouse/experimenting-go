package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	odd := oddSum(sum, numbers...)
	even := evenSum(sum, numbers...)
	fmt.Println(odd)
	fmt.Println(even)

	moreNub := []int{92, 238, 381, 382, 490, 231, 3829, 123, 491, 390, 232}
	even2, odd2 := oddAndEven(sum, moreNub...)
	fmt.Println(odd2)
	fmt.Println(even2)
}

func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func evenSum(f func(x ...int) int, y ...int) int {
	var even []int
	for _, v := range y {
		if v%2 == 0 {
			even = append(even, v)
		}
	}
	return f(even...)
}

func oddSum(f func(x ...int) int, y ...int) int {
	var odd []int
	for _, v := range y {
		if v%2 != 0 {
			odd = append(odd, v)
		}
	}
	return f(odd...)
}

func oddAndEven(f func(x ...int) int, z ...int) (int, int) {
	var odd, even []int
	for _, v := range z {
		if v%2 == 0 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}
	return f(even...), f(odd...)
}
