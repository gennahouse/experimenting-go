package main

import (
	"fmt"
)

func main() {
	a := 10
	b := "Hello"
	c := 10.3020
	fmt.Println(a, b, c)
	a = 20
	b = "Jimmy"
	c = 90.123654
	fmt.Println(a, b, c)
	//Printing type of those variables
	fmt.Printf("%T\t%T\t%T\n", a, b, c)
}
