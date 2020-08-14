//Read and Write using package os
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	filename := "./examples/some_text.txt"
	file, _ := os.Open(filename)

	defer file.Close() // This will be invoked before exit from main

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	out, err := os.Create("output.txt")

	if err != nil {
		log.Fatalln("Cant create file: ", err)
	}

	defer out.Close()

	for _, str := range []string{"Microsoft", "Apple", "Google", "Go!"} {
		bytes, err := out.WriteString(str)
		if err != nil {
			log.Fatalln("Error writing to file: ", err)
		}
		fmt.Printf("Wrote %d bytes to file\n", bytes)
	}
}
