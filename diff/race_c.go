package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("gimve me a number")
		return
	}
	numGR, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(nil)
		return
	}
	var w sync.WaitGroup
	var i int
	var m sync.Mutex
	k := make(map[int]int)
	k[1] = 12
	for i = 0; i < numGR; i++ {
		w.Add(1)
		go func(j int) {
			defer w.Done()
			m.Lock()
			k[j] = j
			m.Unlock()
		}(i)
	}

	w.Wait()
	fmt.Printf("k = %v\n", k)
}
