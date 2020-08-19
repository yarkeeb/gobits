package main

import (
	"fmt"
	"time"
)

func writeToChanel(c chan int, x int) {
	fmt.Println(x)
	c <- x
	close(c)
	fmt.Println(x)
}

func main() {
	c := make(chan int)
	go writeToChanel(c, 10)
	time.Sleep(time.Second)
}
