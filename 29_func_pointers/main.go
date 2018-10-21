package main

import (
	"fmt"
)

func main() {
	x := 50
	y := &x
	changePoinVal(y)

	fmt.Printf("x: %T\ty:%T\n", x, y)
	fmt.Println(x, *y)

}

func changePoinVal(i *int) {
	*i = 50
	fmt.Println("value of both x,y is", *i)
}
