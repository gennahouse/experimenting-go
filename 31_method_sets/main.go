package main

import (
	"fmt"
)

type person struct {
	first string
}

func main() {
	mario := person{
		first: "Mario",
	}
	fmt.Println(mario)
	changeMe(&mario)
	fmt.Println(mario)
}

func changeMe(p *person) {
	(*p).first = "Mary"
}
