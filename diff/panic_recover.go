package main

import (
	"fmt"
)

func a() {
	fmt.Println("Inside a()")
	defer func() {
		if c:= recover(); c != nil {
			fmt.Println("Recover inside a()")
		}
	}()
	fmt.Println("About to call b()")
	b()
	fmt.Println("b() exited")
	fmt.Println("exiting a()")
}

func b() {
	fmt.Println("inside b()")
	panic("Panic in b()")
	fmt.Println("exiting b()")
}

func main() {
	a()
	fmt.Println("mail() ended")
}