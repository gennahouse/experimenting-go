package main

import "fmt"

func main() {
	phoneNums := map[string]int{
		"Gennaro": 3289078276,
		"Hellen":  44074892817,
		"karol":   90287372873,
	}

	fmt.Println(phoneNums["Gennaro"])
	phoneNums["Gennaro"] = 3278909123
	fmt.Printf("%T\n", phoneNums["Gennaro"])
	fmt.Println(phoneNums["Gennaro"])

	x := make([]int, 5, 10)
	x = append(x[:0], 10, 20, 30, 40, 50)
	fmt.Println(x)
	x = append(x[:1], x[4:]...)
	fmt.Println(x)
	y := make([][]int, 10, 10)

	y[1] = x
	fmt.Println(y)
	y[2] = append(y[2], 10, 20, 30, 40)
	fmt.Println(y)

	c := make(map[int]int)

	c[1] = 100

	fmt.Println(c)

}
