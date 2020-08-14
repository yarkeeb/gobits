package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		panic("Not enough args")
	}

	fmt.Println("thanks, args received")
}
