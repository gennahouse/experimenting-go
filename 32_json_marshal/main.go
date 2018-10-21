package main

import (
	"encoding/json"
	"fmt"
)

type order struct {
	ID        uint
	Pair      string
	Amount    float64
	Direction string
}

func main() {

	o1 := order{
		ID:        1,
		Pair:      "Btc/Usd",
		Amount:    0.07835000,
		Direction: "Buy",
	}
	o2 := order{
		ID:        2,
		Pair:      "Btc/Eth",
		Amount:    0.90217000,
		Direction: "Buy",
	}
	o3 := order{
		ID:        3,
		Pair:      "Btc/Trx",
		Amount:    1.12170000,
		Direction: "Sell",
	}

	orders := []order{
		o1,
		o2,
		o3,
	}
	bs, err := json.Marshal(orders)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))

}
