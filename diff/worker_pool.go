package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

type Client struct {
	id      int
	integer int
}

type Data struct {
	job    Client
	square int
}

var (
	size    = 10
	clients = make(chan Client, size)
	data    = make(chan Data, size)
)

func worker(w *sync.WaitGroup) {
	for c := range clients {
		square := c.integer * c.integer
		output := Data{c, square}
		data <- output
		time.Sleep(time.Second)
	}
	w.Done()
}

func makeWP(n int) {
	var w sync.WaitGroup
	for i := 0; i < n; i++ {
		w.Add(1)
		go worker(&w)
	}
	w.Wait()
	close(data)
}

func create(n int) {
	for i := 0; i < n; i++ {
		c := Client{i, i}
		clients <- c
	}
	close(clients)
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("need num jobs and num workers")
		os.Exit(1)
	}

	jobs, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	workers, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	go create(jobs)
	finished := make(chan interface{})
	go func() {
		for d := range data {
			fmt.Printf("Client id: %d\tint: ", d.job.id)
			fmt.Printf("%d\tSquare: %d\n", d.job.integer, d.square)
		}
		finished <- true
	}()
	makeWP(workers)
	fmt.Printf(": %v\n", <-finished)

}
