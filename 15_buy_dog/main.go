package main

import (
	"fmt"
)

type dog struct {
	name, surname string
	age           int
}

type pug struct {
	dog
	price float64
}

func (p pug) callPug() {
	fmt.Println("Hey ", p.surname, p.name)
}

func main() {
	balance := 1000.00
	newBalance, myPug := buyPug(balance)
	myPug.callPug()
	balance = newBalance

	x := struct {
		money float64
	}{
		money: balance,
	}

	fmt.Println(x.money)

}

func buyPug(b float64) (float64, pug) {
	if b >= 843.80 {
		b -= 843.80
		p := pug{
			dog: dog{
				name:    "Anthony The Second",
				surname: "Sir",
				age:     1,
			},
			price: 843.80,
		}
		fmt.Println("Pug purchesed correctly")
		return b, p
	} else {
		fmt.Println("your funds are not enough")
	}
	return 0, pug{}
}
