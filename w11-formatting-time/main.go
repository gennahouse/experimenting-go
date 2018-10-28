package main

import (
	"fmt"
	"time"
)

func main() {
	timeNow := time.Now()
	fmt.Println(timeNow.Format("02-01-2006"))

	timen := time.Now().Format("01-02-2006")
	fmt.Println("American data:", timen)

	timex := time.Now().Format("03:04:05")
	fmt.Println(timex)
}
