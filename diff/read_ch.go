package main

import (
	"fmt"
	"time"
)

func writeToChannel(c chan int, x int) {
	fmt.Println(x)
	c <- x
	close(c)
	fmt.Println(x)
}

func main() {
	c := make(chan int)
	go writeToChannel(c, 10)
	time.Sleep(1 * time.Second)
	fmt.Println("Read:", <-c)
	time.Sleep(1 * time.Second)

	_, ok := <-c
	if ok {
		fmt.Println("chan in open")
	} else {
		fmt.Println("chan is closed")
	}

}
