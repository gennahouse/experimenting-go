package main

import (
	"fmt"
)

//vars scope package: typed vars
//This is not how you should declere vars that work only in a func scope

var a string
var b int
var c float64
var d byte //alias unit8

func main() {
	a = "Hello"
	b = 12
	c = 3.14
	d = 72
	//This is how you print multiple variables using Println
	fmt.Println(a, b, c, d)
}
