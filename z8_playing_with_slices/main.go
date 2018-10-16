package main

import (
	"fmt"
)

func main() {
	s := []int{
		10,
		20,
		30,
		40,
		50,
	}

	for i, v := range s {
		fmt.Println(i, v)
	}
	fmt.Println(s[:3])
	fmt.Println(s[2:])
	fmt.Println(s[1:4])
	s = append(s, 60)
	fmt.Println(s)
	s = append(s[:3])
	fmt.Println(s)
	s = append(s[:1])
	fmt.Println(s)

	x := make([]int, 1, 1)
	x[0] = 100
	fmt.Println(x)

	z := [][]int{
		s,
		x,
	}

	fmt.Println(z)
	 
	z[1][0] = 79
	fmt.Println(z)

}
