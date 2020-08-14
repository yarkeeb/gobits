package main

import (
	"errors"
	"fmt"
)

func returnError(a, b int) error {
	if a == b {
		err := errors.New("Error in returnError() func")
		return err
	}
	return nil
}

func main() {
	err := returnError(1, 2)
	if err == nil {
		fmt.Println("returnError() returned normally")
	} else {
		fmt.Println(err)
	}

	err = returnError(10, 10)
	if err == nil {
		fmt.Println("returnError() returned normally")
	} else {
		fmt.Println(err)
	}

	if err.Error() == "Error in returnError() func" {
		fmt.Println("!!")
	}
}