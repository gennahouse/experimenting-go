package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c circle) area() float64 {
	area := math.Pi * c.radius * c.radius
	return area
}

func (s square) area() float64 {
	area := s.side * s.side
	return area
}

func info(s shape) {
	x := s.area()
	fmt.Println(x)
}
func main() {

	s := square{
		side: 10,
	}
	c := circle{
		radius: 10.0,
	}

	info(s)
	info(c)

}
