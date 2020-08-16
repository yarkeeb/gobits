package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Tree struct {
	Value int
	Left  *Tree
	Right *Tree
}

func traverse(t *Tree) {
	if t == nil {
		return
	}
	traverse(t.Left)
	fmt.Print(t.Value, " ")
	traverse(t.Right)
}

func create(n int) *Tree {
	var t *Tree
	rand.Seed(time.Now().Unix())
	for i := 0; i < 2*n; i++ {
		temp := rand.Intn(n * 2)
		t = insert(t, temp)
	}
	return t
}

func insert(t *Tree, n int) *Tree {
	if t == nil {
		return &Tree{n, nil, nil}
	}
	if n == t.Value {
		return t
	}
	if n < t.Value {
		t.Left = insert(t.Left, n)
		return t
	}
	t.Right = insert(t.Right, n)
	return t
}

func main() {
	tree := create(10)
	fmt.Println("root val", tree.Value)
	traverse(tree)
	fmt.Println()
	tree = insert(tree, -10)
	tree = insert(tree, -2)
	traverse(tree)
	fmt.Println()
	fmt.Println("rot val", tree.Value)
}
