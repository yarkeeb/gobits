//Work with some func's from strings package

package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello world"
	s_equal := "Hello world"
	s_not := "world, Hello"

	// Simply iterate over string
	for _, c := range s {
		fmt.Printf("%c\n", c)
	}

	arr := []rune(s)

	for _, r := range arr {
		fmt.Printf("%d\n", r)
	}

	fmt.Printf("s := %s\n", s)
	fmt.Printf("s_equal := %s\n", s_equal)
	fmt.Printf("s_not := %s\n", s_not)

	//Compare returns an integer comparing two strings lexicographically.
	// The result will be 0 if a==b, -1 if a < b, and +1 if a > b
	fmt.Printf("equal: %d\n", strings.Compare(s, s_equal))
	fmt.Printf("not equal: %d\n", strings.Compare(s, s_not))

	//Contains reports whether substr is within s.
	fmt.Printf("check if \"world\" is substr of s: %v\n", strings.Contains(s, "world"))
	fmt.Printf("check in \"cat\" is substr of s: %v\n", strings.Contains(s, "cat"))

	// ContainsAny reports whether any Unicode code points in chars are within s.
	fmt.Printf("any of \"abcd\" in s: %v\n", strings.ContainsAny(s, "abcd"))
	fmt.Printf("any of \"mkt\" in s: %v\n", strings.ContainsAny(s, "mkt"))

	// ContainsRune reports whether the Unicode code point r is within s.
	fmt.Printf("find rune 72 is s: %v\n", strings.ContainsRune(s, 72))

	//ToLower
	fmt.Printf("lowercase: %s\n", strings.ToLower(s))

	//ToUpper
	fmt.Printf("uppercase: %s\n", strings.ToUpper(s))

	sentence := "Simple sentence with many words"

	words := strings.Split(sentence, " ")

	fmt.Printf("%v\n", words)

	//Split with many whitespaces : strings.Fields

	words = strings.Fields(sentence)

	fmt.Printf("%v\n", words)

}
