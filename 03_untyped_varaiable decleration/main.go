package main

import (
	"fmt"
)

//untyped package scope variable declered
//They should be declered inside the func instead
var a = 10
var b = "Hello"
var c = 10.10

func main() {
	fmt.Println(a, b, c)
}
