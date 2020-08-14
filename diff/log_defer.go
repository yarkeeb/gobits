package main

import (
	"fmt"
	"log"
	"os"
)

//LOGFILE is log file
var LOGFILE = "/tmp/mGo.log"

func one(logger *log.Logger) {
	logger.Println("-- FUNCTION one --")
	defer logger.Println("-- FUNCTION one --")

	for i := 0; i < 10; i++ {
		logger.Println(i)
	}
}

func two(logger *log.Logger) {
	logger.Println("-- FUNCTION two --")
	defer logger.Println("-- FUNCTION two --")

	for i := 0; i < 10; i++ {
		logger.Println(i)
	}
}

func main() {
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	logger := log.New(f, "logDefer", log.LstdFlags)
	logger.Println("hello there")
	logger.Println("another log entry")

	one(logger)
	two(logger)
}