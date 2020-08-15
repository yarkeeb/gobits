package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 100; i++ {
		if i%20 == 0 {
			continue
		}
		if i == 95 {
			break
		}
		fmt.Print(i, " ")
	}

	fmt.Println()
	i := 10
	for {
		if i < 0 {
			break
		}
		fmt.Print(i, " ")
		i--
	}
	fmt.Println()

	i = 0
	expression := true
	for ok := true; ok; ok = expression {
		if i > 10 {
			expression = false
		}
		fmt.Print(i, " ")
		i++
	}
	fmt.Println()

	arr := [5]int{0, 1, -1, 2, -2}
	for i, v := range arr {
		fmt.Println("index:", i, "value:", v)
	}
}
