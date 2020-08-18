package main

import (
	"fmt"
	"os"
	"reflect"
)

type a struct {
	X int
	Y float64
	Z string
}

type b struct {
	F int
	G int
	H string
	I float64
}

func main() {
	x := 100
	xRefl := reflect.ValueOf(&x).Elem()
	xType := xRefl.Type()
	fmt.Printf("type of x is %s.\n", xType)

	A := a{100, 200.12, "struct a"}
	B := b{1, 2, "struct b", -1.2}
	var r reflect.Value

	arguments := os.Args
	if len(arguments) == 1 {
		r = reflect.ValueOf(&A).Elem()
	} else {
		r = reflect.ValueOf(&B).Elem()
	}

	iType := r.Type()
	fmt.Printf("i type: %s\n", iType)
	fmt.Printf("The %d fieds of %s are:\n", r.NumField(), iType)
	for i := 0; i < r.NumField(); i++ {
		fmt.Printf("field name: %s ", iType.Field(i).Name)
		fmt.Printf("with type: %s ", r.Field(i).Type())
		fmt.Printf("and value: %v\n", r.Field(i).Interface())
	}
}
