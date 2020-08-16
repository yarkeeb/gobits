package main

import "fmt"

const SIZE = 15

type Node struct {
	Value int
	Next  *Node
}

type HashTable struct {
	Table map[int]*Node
	Size  int
}

func hashFunction(i, size int) int {
	return (i % size)
}

func insert(table *HashTable, value int) int {
	index := hashFunction(value, table.Size)
	e := Node{Value: value, Next: table.Table[index]}
	table.Table[index] = &e
	return index
}

func traverse(table *HashTable) {
	for k := range table.Table {
		if table.Table[k] != nil {
			t := table.Table[k]
			for t != nil {
				fmt.Printf("%d -> ", t.Value)
				t = t.Next
			}
			fmt.Println()
		}
	}
}

func main() {
	table := make(map[int]*Node, SIZE)
	hash := &HashTable{Table: table, Size: SIZE}
	fmt.Println("Num of spaces", hash.Size)
	for i := 0; i < 120; i++ {
		insert(hash, i)
	}
	traverse(hash)
}
