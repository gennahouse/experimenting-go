package main

import (
	"fmt"
)

func main() {
	tex := map[int]int{
		1: 20,
		2: 40,
		3: 50,
	}
	fmt.Println(tex)
	delete(tex, 1)
	tex[1], tex[4] = 10, 20
	fmt.Println(tex)
	if v, ok := tex[10]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("No such value")
	}
	for k, va := range tex {
		fmt.Println(k, va)
	}
}
