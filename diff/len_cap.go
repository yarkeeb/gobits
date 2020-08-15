package main

import (
	"fmt"
)

func printSlice(x []int) {
	for _, n := range x {
		fmt.Print(n, " ")
	}
	fmt.Println()
}

func main() {
	s := []int{-1, 0, 4}
	fmt.Printf("s: ")
	printSlice(s)
	fmt.Printf("Cap: %d, Len: %d\n", cap(s), len(s))
	s = append(s, -100)
	fmt.Printf("s: ")
	printSlice(s)
	fmt.Printf("Cap: %d, Len: %d\n", cap(s), len(s))

	s = append(s, -2)
	s = append(s, -3)
	s = append(s, -4)
	printSlice(s)
	fmt.Printf("Cap: %d, Len: %d\n", cap(s), len(s))
}
