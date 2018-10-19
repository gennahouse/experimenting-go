package main

import (
	"fmt"
)

func main() {
	func() {
		x := 10
		x *= 10
		fmt.Println(x)
	}()
}
