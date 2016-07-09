package main

import (
	"fmt"
	"go/build"
)

func main() {
	fmt.Println(build.Default.GOPATH)
}
