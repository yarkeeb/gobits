package main

import (
	"fmt"
)

type square struct {
	X float64
}

type circle struct {
	R float64
}

type rectangle struct {
	X float64
	Y float64
}

func tellInterface(x interface{}) {
	switch v := x.(type) {
	case square:
		fmt.Println("square")
	case circle:
		fmt.Printf("%v is a circle\n", v)
	case rectangle:
		fmt.Println("rectangle")
	default:
		fmt.Printf("unknown %T\n", v)
	}
}

func main() {
	x := circle{R: 10}
	tellInterface(x)
	y := rectangle{X: 4, Y: 1}
	tellInterface(y)
	z := square{X: 10}
	tellInterface(z)
	tellInterface(10)
}
