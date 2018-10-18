package main

import (
	"fmt"
)

func main() {
	italy := map[string]int{
		"Torino":  1908721,
		"Roma":    38928172,
		"Firenze": 89098,
	}
	fmt.Println(italy)
	delete(italy, "Firenze")

	if _, ok := italy["Roma"]; ok {
		fmt.Println("Roma is a valid city")
	}
	if _, ok := italy["Firenze"]; ok {
		fmt.Println("Firenze is a valid city")
	} else {
		fmt.Println("Firenze is not a valid city")
	}

	uk := make(map[int]int)
	uk[1] = 12328912
	uk[2] = 448934958

	fmt.Println(uk)

	for i, v := range uk {
		fmt.Println(i, v)
	}
}
