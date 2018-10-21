package main

import (
	"encoding/json"
	"fmt"
)

type order struct {
	ID        uint    `json="ID"`
	Pair      string  `json="Pair"`
	Amount    float64 `json="Amount"`
	Direction string  `json="Direction"`
}

func main() {
	s := `[{"ID":1,"Pair":"Btc/Usd","Amount":0.07835,"Direction":"Buy"},{"ID":2,"Pair":"Btc/Eth","Amount":0.90217,"Direction":"Buy"},{"ID":3,"Pair":"Btc/Trx","Amount":1.1217,"Direction":"Sell"}]`

	bs := []byte(s)

	var orders []order

	err := json.Unmarshal(bs, &orders)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range orders {
		fmt.Println("ID:", v.ID, "Pair:", v.Pair, "Amount:", v.Amount, "Direction:", v.Direction)
	}

}
