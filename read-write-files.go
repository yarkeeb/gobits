//Examples of reading and writing files in Go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	filename := "./examples/some_text.txt"

	content, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatalln("Error reading file", filename)
	}

	fmt.Println(string(content))

	output := "output.txt"

	err = ioutil.WriteFile(output, content, 0644)

	if err != nil {
		log.Fatalln("Error writing file", filename)
	}
}
