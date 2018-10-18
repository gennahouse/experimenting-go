package main

import (
	"fmt"
)

func main() {
	total := 1
	for i := 1; i <= 4; i++ {
		total *= i
	}
	fmt.Println(total)

	fuLo := factLoop(4)
	fmt.Println(fuLo)
}

func factLoop(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}
