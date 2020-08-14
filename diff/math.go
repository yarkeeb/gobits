package main

import (
	"fmt"
	"os"
	"strconv"
)


func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please give one or more floats")
		os.Exit(1)
	}
	arguments := os.Args
	sum := 0.0
	numLen := 0

	for i := 1; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)
		if err == nil {
			sum += n
			numLen++
		}
	}

	med := sum / float64(numLen)
	fmt.Printf("sum: %f, med: %f\n", sum, med)

}