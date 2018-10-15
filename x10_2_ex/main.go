package main

import (
	"fmt"
)

func main() {
	x := []int{42, 43, 44, 45, 46}
	fmt.Println(x)
	x = append(x[:0], 47, 48, 49, 50, 51)
	fmt.Println(x)
	x = append(x[:0], 44, 45, 46, 47, 48)
	fmt.Println(x)
	x = append(x[:0], 43, 44, 45, 46, 47)
	fmt.Println(x)

	y := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	t := []int{}
	a := append(t, y[:5]...)
	fmt.Println(a)
	b := append(t, y[5:]...)
	fmt.Println(b)
	c := append(t, y[2:7]...)
	fmt.Println(c)
	d := append(t, y[1:6]...)
	fmt.Println(d)
}
