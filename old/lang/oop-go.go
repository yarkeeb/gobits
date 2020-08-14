//Some OOP
package main

import (
	"fmt"
)

type Cat struct {
	name string
	age  int
}

func (cat Cat) Meow() string {
	return cat.name + " did meow!"
}

func (cat *Cat) Grow(age int) {
	cat.age += age
	fmt.Printf("%v levels up by %d years up to %d!\n", cat.name, age, cat.age)
}

func main() {
	Vaider := Cat{name: "Vaider the cat", age: 15}

	fmt.Printf("%v\n", Vaider.Meow())

	Vaider.Grow(10)
}
