package main

import (
	"fmt"
)

func main() {
	fmt.Println("string is collection on runes:", []byte("Test string"))
	str := []byte("Test string")
	for x, y := range str {
		fmt.Println(x, y)
		fmt.Printf("Char %c\n", str[x])
	}
	fmt.Printf("%s\n", str)
}
