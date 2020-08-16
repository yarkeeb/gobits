package main

import (
	"fmt"
	s "strings"
)

var f = fmt.Printf

func main() {
	upper := s.ToUpper("Hello There!")
	f("To upper: %s\n", upper)
	f("To lower: %s\n", s.ToLower(upper))
	f("%s\n", s.Title("ThIs WilL bE a tItLe"))
	f("Equal fold: %v\n", s.EqualFold("Test string", "TEST STRING"))
	f("Equal fold: %v\n", s.EqualFold("Test string", "TEST STR"))

	f("Fields: %v\n", s.Fields("This is a string!"))
	f("%s\n", s.Split("abc def", " "))
	f("%s\n", s.Replace("abcd ef", "", "_", -1))
	f("%s\n", s.Replace("abcd ef", "", "_", 4))
	f("%s\n", s.Replace("abcd ef", "", "_", 2))

	lines := []string{"line 1", "line 2", "line 3"}
	f("Join: %s\n", s.Join(lines, "++"))

	f("SplitAfter: %s\n", s.SplitAfter("123+++432++", "++"))
}
