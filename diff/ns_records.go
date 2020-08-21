package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("provide hostname")
		return
	}

	domain := os.Args[1]
	NSs, err := net.LookupNS(domain)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, NS := range NSs {
		fmt.Println(NS.Host)
	}
}
