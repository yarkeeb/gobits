package main

import (
	"io"
	"os"
)

func main() {
	myString := ""
	arguments := os.Args

	if len(arguments) == 1 {
		myString = "Please give me at least one argument"
	} else {
		myString = arguments[1]
	}

	io.WriteString(os.Stdout, "This is standart output\n")
	io.WriteString(os.Stderr, myString)
	io.WriteString(os.Stderr, "\n")
}
