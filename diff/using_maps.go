package main

import (
	"fmt"
)

func main() {
	iMap := make(map[string]int)
	iMap["k1"] = 1
	iMap["k2"] = 2
	fmt.Println("iMap:", iMap)

	anotherMap := map[string]int{
		"k1": 12,
		"k2": 13,
	}
	fmt.Println("anotherMap:", anotherMap)

	delete(anotherMap, "k1")
	delete(anotherMap, "k1")
	delete(anotherMap, "k1")
	fmt.Println("anotherMap:", anotherMap)

	_, ok := iMap["doesitexist"]
	if ok {
		fmt.Println("exists")
	} else {
		fmt.Println("does NOT exist")
	}

	for key, value := range iMap {
		fmt.Println(key, value)
	}
}
