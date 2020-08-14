package main

import "C"

import (
	"fmt"
)

//export PrintMessage
func PrintMessage() {
	fmt.Println("A Go func")
}

//export Multiply
func Multiply(a, b int) int {
	return a * b
}

func main() {
}
