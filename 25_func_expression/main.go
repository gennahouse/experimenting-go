package main

import (
	"fmt"
)

func main() {
	x := func() {
		fmt.Println("Hello from func")
	}
	x()
}
