package main

import "fmt"

func main() {
	x := 10
	for {
		if x >= 15 {
			break
		}
		fmt.Println(x, "< 15")
		x++
	}
}
