package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var Password = secret{password: "mypassword"}

type secret struct {
	RWM      sync.RWMutex
	M        sync.Mutex
	password string
}

func Change(c *secret, pass string) {
	c.RWM.Lock()
	fmt.Println("LChange")
	time.Sleep(10 * time.Second)
	c.password = pass
	c.RWM.Unlock()
}

func show(c *secret) string {
	c.RWM.RLock()
	fmt.Print("Show")
	time.Sleep(3 * time.Second)
	defer c.RWM.RUnlock()
	return c.password
}

func showWithLock(c *secret) string {
	c.M.Lock()
	fmt.Println("showWithLock")
	time.Sleep(3 * time.Second)
	defer c.M.Unlock()
	return c.password
}

func main() {
	var showFunction = func(c *secret) string { return "" }
	if len(os.Args) != 2 {
		fmt.Println("using sync.RWMutex")
		showFunction = show
	} else {
		showFunction = showWithLock
		fmt.Println("using sync.Mutex")
	}

	var waitGroup sync.WaitGroup
	for i := 0; i < 15; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			fmt.Println("pass:", showFunction(&Password))
		}()
	}

	go func() {
		waitGroup.Add(1)
		defer waitGroup.Done()
		Change(&Password, "123456")
	}()

	waitGroup.Wait()
	fmt.Println("pass:", showFunction(&Password))
}
