package main

import (
	"fmt"
	"math"

	"./myinterface"
)

type square struct {
	X float64
}

type circle struct {
	R float64
}

func (s square) Area() float64 {
	return s.X * s.X
}

func (s square) Perimeter() float64 {
	return s.X * 4
}

func (c circle) Area() float64 {
	return c.R * c.R * math.Pi
}

func (c circle) Perimeter() float64 {
	return 2 * c.R * math.Pi
}

func Calculate(x myinterface.Shape) {
	_, ok := x.(circle)
	if ok {
		fmt.Println("circle")
	}

	v, ok := x.(square)
	if ok {
		fmt.Println("square:", v)
	}

	fmt.Println(x.Area())
	fmt.Println(x.Perimeter())
}

func main() {
	x := square{X: 10}
	fmt.Println("Perimeter:", x.Perimeter())
	Calculate(x)
	c := circle{R: 10}
	Calculate(c)
}
