package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s)
	integers := make([]int, 2)
	fmt.Println(integers)
	integers = nil
	fmt.Println(integers)

	arr := [5]int{-1, -2, -3, -4, -5}
	refarr := arr[:]

	fmt.Println(arr)
	fmt.Println(refarr)
	arr[4] = -100
	fmt.Println(refarr)

	sb := make([]byte, 5)
	fmt.Println(sb)
	twoD := make([][]int, 3)
	fmt.Println(twoD)
	fmt.Println()
	for i := 0; i < len(twoD); i++ {
		for j := 0; j < 2; j++ {
			twoD[i] = append(twoD[i], i*j)
		}
	}

	for _, x := range twoD {
		for i, y := range x {
			fmt.Println("i:", i, "value", y)
		}
	}
	fmt.Println()

}
