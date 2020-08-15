package main

import (
	"fmt"
	"sort"
)

type structure struct {
	person string
	height int
	weight int
}

func main() {
	mySlice := make([]structure, 0)
	mySlice = append(mySlice, structure{"John", 180, 90})
	mySlice = append(mySlice, structure{"Bill", 134, 54})
	mySlice = append(mySlice, structure{"Willy", 150, 45})
	mySlice = append(mySlice, structure{"Pepe", 145, 80})
	mySlice = append(mySlice, structure{"PekMa", 176, 75})

	fmt.Println("0:", mySlice)

	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].height < mySlice[j].height
	})
	fmt.Println("<:", mySlice)
	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].height > mySlice[j].height
	})
	fmt.Println(">:", mySlice)

}
