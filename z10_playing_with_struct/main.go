package main

import (
	"fmt"
)

type person struct {
	first  string
	second string
	age    int
}

type agent struct {
	person
	ltk bool
}

func main() {

	jb := agent{
		person: person{
			first:  "James",
			second: "Bond",
			age:    32,
		},
		ltk: true,
	}

	fmt.Println(jb)

	fmt.Println(jb.first, jb.second, jb.age, jb.ltk)

	mp := struct {
		name    string
		surname string
		age     int
		gender  string
	}{
		name:    "Miss",
		surname: "Moneypenny",
		age:     24,
		gender:  "female",
	}

	fmt.Println(mp)

	fmt.Println(mp.name, mp.surname, mp.age, mp.gender)

}
