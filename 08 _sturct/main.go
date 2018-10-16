package main

import (
	"fmt"
)

type person struct {
	firstName   string
	lastName    string
	favIceCream []string
}

func main() {

	jack := person{
		firstName: "Jack",
		lastName:  "Gopher",
		favIceCream: []string{
			"Vanilla",
			"choccolate",
		},
	}

	fmt.Println(jack)

	mary := person{
		firstName: "Mary",
		lastName:  "Loo",
		favIceCream: []string{
			"Cigaratte Flavor LOL",
			"Banana",
		},
	}

	fmt.Println(mary)

	for i, v := range jack.favIceCream {
		fmt.Println(i, v)
	}

	for j, x := range mary.favIceCream {
		fmt.Println(j, x)
	}
}
