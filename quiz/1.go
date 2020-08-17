package main

import (
	"fmt"
	"reflect"
)

type T struct {
	X int
	x int
}

func (T) M() {}
func (T) m() {}

func main() {
	v := reflect.ValueOf(T{})
	fmt.Println(v.NumField(), v.NumMethod())
}
