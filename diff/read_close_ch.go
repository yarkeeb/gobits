package main

import "fmt"

func main() {
	willClose := make(chan int, 10)
	willClose <- 1
	willClose <- 0
	willClose <- 2

	<-willClose
	<-willClose
	<-willClose

	close(willClose)
	read := <-willClose
	fmt.Println(read)

}
