package main

import (
	"fmt"
)

func main() {
	defer delayed()
	fmt.Println("This is not using defer")
}

func delayed() {
	defer func() {
		fmt.Println("This is using defer inside func")
	}()
	fmt.Println("This would have been printed for last if there wasn't defer")
}
