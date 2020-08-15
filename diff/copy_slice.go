package main

import (
	"fmt"
)

func main() {
	a6 := []int{-10, 1, 2, 3, 4, 5}
	a4 := []int{-1, -2, -3, -4}
	fmt.Println("a6:", a6)
	fmt.Println("a4:", a4)
	copy(a6, a4)
	fmt.Println("a6:", a6)
	fmt.Println("a4:", a4)
	fmt.Println()

	b6 := []int{-10, 1, 2, 3, 4, 5}
	b4 := []int{-1, -2, -3, -4}
	fmt.Println("b6:", b6)
	fmt.Println("b4:", b4)
	copy(b4, b6)
	fmt.Println("b6:", b6)
	fmt.Println("b4:", b4)
	fmt.Println()

	arr4 := [4]int{4, -4, 4, -4}
	s6 := []int{1, -1, 2, -2, 3, -3}
	copy(s6, arr4[0:])
	fmt.Println("arr4:", arr4)
	fmt.Println("s6:", s6)
	fmt.Println()

	arr5 := [5]int{5, 5, 5, 5, 5}
	s7 := []int{7, 7, 7, 7, 7, 7, 7}
	copy(arr5[0:], s7)
	fmt.Println("arr5:", arr5)
	fmt.Println("s7:", s7)

}
