package main

import (
	"fmt"
)

type car struct {
	color    string
	wheels   int
	maxSpeed float64
	doors    int
	seats    int
}

type lamborghini struct {
	car
	model string
}

type fiat struct {
	car
	model string
}

type airplan struct {
	wings    int
	maxSpeed float64
	seats    int
	color    string
}

type boing struct {
	airplan
	model string
}

type vehicle interface {
	driveDistance(i int)
}

func main() {
	aventador := lamborghini{
		car: car{
			color:    "yellow",
			wheels:   4,
			maxSpeed: 320,
			doors:    2,
			seats:    2,
		},
		model: "aventador",
	}

	punto := fiat{
		car: car{
			color:    "white",
			wheels:   4,
			maxSpeed: 175,
			doors:    4,
			seats:    5,
		},
		model: "punto",
	}

	boing747 := boing{
		airplan: airplan{
			wings:    2,
			maxSpeed: 1300.90,
			seats:    280,
			color:    "blue",
		},
		model: "747",
	}
	whatVehicle(aventador, 2480)
	whatVehicle(punto, 2480)
	whatVehicle(boing747, 2480)
}

func (l lamborghini) driveDistance(i int) {
	driveTime := float64(i) / l.maxSpeed
	fmt.Println("A", l.model, "Takes:", driveTime, "to get at the destination.")
}

func (f fiat) driveDistance(i int) {
	driveTime := float64(i) / f.maxSpeed
	fmt.Println("A", f.model, "Takes:", driveTime, "to get at the destination.")
}

func (b boing) driveDistance(i int) {
	driveTime := float64(i) / b.maxSpeed
	fmt.Println("A", b.model, "Takes:", driveTime, "to get at the destination.")
}

func whatVehicle(v vehicle, i int) {
	switch v.(type) {
	case lamborghini:
		v.(lamborghini).driveDistance(i)
	case fiat:
		v.(fiat).driveDistance(i)
	case boing:
		v.(boing).driveDistance(i)
	default:
		fmt.Println("Not a valid vehicle selected")
	}
	fmt.Printf("%v\n\n", v)

}
