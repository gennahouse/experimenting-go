package main

import (
	"fmt"
)

func main() {
	x := retFun()
	x()
}

func retFun() func() {
	return func() {
		fmt.Println("Hello from the returned function!")
	}
}
