package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.BindEnv("GOMAXPROCS")
	val := viper.Get("GOMAXPROCS")
	fmt.Println(val)
	viper.Set("GOMAXPROCS", 10)
	val = viper.Get("GOMAXPROCS")
	fmt.Println("GOMAXPROCS", val)

	viper.BindEnv("NEW_VAR")
	val = viper.Get("NEW_VAR")
	if val == nil {
		fmt.Println("NEW_VAR not defined")
		return
	}
	fmt.Println(val)
}
