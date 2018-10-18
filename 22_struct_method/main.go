package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func main() {
	jhonny := person{
		first: "Jhon",
		last:  "Deacon",
		age:   68,
	}
	jhonny.speak()
}

func (p person) speak() {
	fmt.Println(p.first, p.last, p.age)
}
