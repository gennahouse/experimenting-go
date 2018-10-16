package main

import (
	"fmt"
)

func main() {
	j := struct {
		name string
		last string
		m    map[int]int
		s    []int
	}{
		name: "jack",
		last: "nonsense",
		m: map[int]int{
			1: 10,
			2: 20,
			3: 20,
		},
		s: []int{
			2,
			4,
			6,
			8,
			10,
		},
	}
	fmt.Println(j)
}
