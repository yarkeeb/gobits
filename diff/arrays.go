package main

import "fmt"

func main() {
	arr := [4]int{1, -1, 2, -2}
	twoD := [4][4]int{{1, 2, 3, 4}, {4, 3, 2, 1}, {5, 4, 3, 2}, {2, 3, 4, 5}}
	threeD := [2][2][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}

	fmt.Println("len of ", arr, "is", len(arr))
	fmt.Println("first ele of", twoD, "is", twoD[0][0])
	fmt.Println("len of", threeD, "is", len(threeD))

	for i := 0; i < len(threeD); i++ {
		v := threeD[i]
		for j := 0; j < len(v); j++ {
			m := v[j]
			for k := 0; k < len(m); k++ {
				fmt.Print(m[k], " ")
			}
		}
		fmt.Println()
	}

	for _, v := range threeD {
		for _, m := range v {
			for _, s := range m {
				fmt.Print(s, " ")
			}
		}
		fmt.Println()
	}
}
