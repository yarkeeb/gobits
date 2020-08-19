package main

import "fmt"

func main() {
	count := 0

	for i := range [256]struct{}{} {
		m, n := byte(i), int8(i)
		if n == -n {
			fmt.Println("n", n, -n)
			count++
		}
		if m == -m {
			fmt.Println("m", m, -m)
			count++
		}
	}
	fmt.Println(count)
}
