//Some functions examples
package main

import (
	"fmt"
)

func multiply(x, y int) int {
	return x * y
}

func modify_parameter(x *int) {
	*x += 5
}

func Say(s string) string {
	return s + " YEAH"
}

func main() {
	fmt.Printf("%v\n", multiply(5, 6))

	x := 5
	fmt.Printf("x before: %v\n", x)
	p := &x
	modify_parameter(p)
	fmt.Printf("x after: %v\n", x)

	fmt.Printf("%v\n", Say("AWHH"))
}
