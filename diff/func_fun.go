package main

import "fmt"

func fun1(i int) int {
	return i + i
}

func fun2(i int) int {
	return i * i
}

func funFun(f func(int) int, v int) int {
	return f(v)
}

func main() {
	fmt.Println("fun1:", funFun(fun1, 1))
	fmt.Println("fun2:", funFun(fun2, 3))
	fmt.Println("Inline:", funFun(func(i int) int { return i * i * i }, 3))
}
