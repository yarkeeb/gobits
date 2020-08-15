package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3}
	a := [3]int{4, 5, 6}

	ref := a[:]
	fmt.Println("existing array:\t", ref)
	t := append(s, ref...)
	fmt.Println("new slice:\t", t)
	s = append(s, ref...)
	fmt.Println("existing slice\t", s)
	s = append(s, s...)
	fmt.Println("s+s:\t\t", s)

}
