package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("You are using ", runtime.Compiler, " ")
	fmt.Println("on a ", runtime.GOARCH, "machine")
	fmt.Println("using go version ", runtime.Version())
	fmt.Println("number of cpus: ", runtime.NumCPU())
	fmt.Println("number of goroutines: ", runtime.NumGoroutine())
}