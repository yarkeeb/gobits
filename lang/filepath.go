package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println("err")
	} else {
		fmt.Println(dir)
	}

}
