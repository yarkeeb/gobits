package main

import (
	"flag"
	"fmt"
	"strings"
)

type NamesFlag struct {
	Names []string
}

func (s *NamesFlag) GetNames() []string {
	return s.Names
}

func (s *NamesFlag) String() string {
	return fmt.Sprint(s.Names)
}

func (s *NamesFlag) Set(v string) error {
	if len(s.Names) > 0 {
		return fmt.Errorf("cannot use names flag more than once")
	}
	names := strings.Split(v, ",")
	for _, item := range names {
		s.Names = append(s.Names, item)
	}
	return nil
}

func main() {
	var manyNames NamesFlag
	minusK := flag.Int("k", 0, "an int")
	minusO := flag.String("o", "test", "name")
	flag.Var(&manyNames, "names", "comma-separated list")

	flag.Parse()

	fmt.Println("k:", *minusK)
	fmt.Println("o:", *minusO)

	for i, item := range manyNames.GetNames() {
		fmt.Println(i, item)
	}

	fmt.Println("remaining:")
	for i, v := range flag.Args() {
		fmt.Println(i, ":", v)
	}
}
