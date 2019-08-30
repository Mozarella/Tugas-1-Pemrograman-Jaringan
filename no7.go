package main

import (
	"fmt"

	"github.com/spf13/viper/"
)

func main() {
	viper.SetConfigFile("./config/env.json")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error Config File!, &5", err)
	}
}
