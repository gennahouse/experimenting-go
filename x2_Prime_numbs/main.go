package main

import (
	"fmt"
)

func main() {
	i := 0
	for {
		i++
		if i >= 100 {
			break
		}
		x := primalityTest(i)

		if x == false {
			continue
		}
		fmt.Println(i, "is a prime number")
	}
}

func primalityTest(n int) bool {
	check := []bool{}
	var valCheck bool
	for i := 2; i < n-1; i++ {
		if n%i != 0 {
			valCheck = true
		} else {
			valCheck = false
		}
		check = append(check, valCheck)
	}
	for _, val := range check {
		if val == false {
			return false
		}
	}
	return true
}
