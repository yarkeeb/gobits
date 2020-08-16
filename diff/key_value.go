package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type element struct {
	Name    string
	Surname string
	ID      string
}

var DATA = make(map[string]element)

func Add(k string, e element) bool {
	if k == "" {
		return false
	}
	if Lookup(k) == nil {
		DATA[k] = e
		return true
	}
	return false
}

func Delete(k string) bool {
	if Lookup(k) != nil {
		delete(DATA, k)
		return true
	}
	return false
}

func Lookup(k string) *element {
	_, ok := DATA[k]
	if ok {
		e := DATA[k]
		return &e
	}
	return nil
}

func Change(k string, e element) bool {
	DATA[k] = e
	return true
}

func Print() {
	for k, d := range DATA {
		fmt.Printf("key: %s, value: %v\n", k, d)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		text = strings.TrimSpace(text)
		tokens := strings.Fields(text)

		switch len(tokens) {
		case 0:
			continue
		case 1:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 2:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 3:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 4:
			tokens = append(tokens, "")
		}

		switch tokens[0] {
		case "PRINT":
			Print()
		case "STOP":
			return
		case "DELETE":
			if !Delete(tokens[1]) {
				fmt.Println("delete failed")
			}
		case "ADD":
			n := element{tokens[2], tokens[3], tokens[4]}
			if !Add(tokens[1], n) {
				fmt.Println("add failed")
			}
		case "LOOKUP":
			n := Lookup(tokens[1])
			if n != nil {
				fmt.Printf("%v\n", n)
			}
		case "CHANGE":
			n := element{tokens[2], tokens[3], tokens[4]}
			if !Change(tokens[1], n) {
				fmt.Println("update failed")
			}
		default:
			fmt.Println("unknown command")
		}
	}
}
