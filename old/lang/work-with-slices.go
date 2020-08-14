// Test some slices functions
package main

import (
	"fmt"
)

func main() {
	letters := []string{"a", "b", "c", "d", "e"}
	fmt.Printf("%v\n", letters)

	var byte_slice []byte

	byte_slice = make([]byte, 5, 5)

	fmt.Printf("%v\n", byte_slice)

	new_slice := make([]byte, 5)

	fmt.Printf("%d, %d\n", len(new_slice), cap(new_slice))

	letters = append(letters, "f")

	fmt.Printf("%v\n", letters)

	fmt.Printf("%v\n", letters[1:4])
	fmt.Printf("%v\n", letters[:5])
	fmt.Printf("%v\n", letters[:])
	fmt.Printf("%v\n", letters[:2:5])

	for index, letter := range letters {
		fmt.Printf("index: %d, value: %v\n", index, letter)
	}
}
