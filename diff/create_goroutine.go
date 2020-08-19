package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	n := flag.Int("n", 10, "number of goroutines")
	flag.Parse()
	count := *n
	fmt.Printf("creating %d goroutines\n", count)

	for i := 0; i < count; i++ {
		go func(x int) {
			fmt.Printf("%d", x)
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println("\nexiting")

}
