// Examples for command-line args usage
package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	str  string
	num  int
	help bool
)

func main() {
	num_args := len(os.Args)

	if num_args < 2 {
		fmt.Println("No args")
	}

	flag.StringVar(&str, "str", "default str", "description")
	flag.IntVar(&num, "num", 5, "description")
	flag.BoolVar(&help, "help", false, "display help")
	flag.Parse()

	if help {
		fmt.Println(">> Display help")
		os.Exit(1)
	}

	fmt.Println(">> Str: ", str)
	fmt.Println(">> Num: ", num)

	args := flag.Args()

	if len(args) > 0 {
		fmt.Println("Flag arg: ", args[0])
	}

}
