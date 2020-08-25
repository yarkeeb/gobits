package main

import (
	"fmt"
	"os"

	"github.com/karrick/godirwalk"
)

func main() {
	dirname := os.Args[1]
	err := godirwalk.Walk(dirname, &godirwalk.Options{
		Callback: func(osPathname string, info *godirwalk.Dirent) error {
			fmt.Printf("%s %s\n", info.ModeType(), osPathname)
			return nil
		},
		ErrorCallback: func(osPathname string, err error) godirwalk.ErrorAction {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			return godirwalk.SkipNode
		},
		Unsorted: true,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
}
