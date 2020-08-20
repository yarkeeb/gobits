package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func timeout(w *sync.WaitGroup, t time.Duration) bool {
	temp := make(chan int)
	go func() {
		defer close(temp)
		time.Sleep(5 * time.Second)
		w.Wait()
	}()

	select {
	case <-temp:
		return false
	case <-time.After(t):
		return true
	}
}

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("need a time duration")
		return
	}

	var w sync.WaitGroup
	w.Add(1)

	t, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	duration := time.Duration(int32(t)) * time.Millisecond
	fmt.Printf("timeout period is %s\n", duration)

	if timeout(&w, duration) {
		fmt.Println("timed out")
	} else {
		fmt.Println("ok")
	}

	w.Done()
	if timeout(&w, duration) {
		fmt.Println("timed out")
	} else {
		fmt.Println("ok")
	}
}
