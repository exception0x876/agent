package main

import (
	"flag"
	"fmt"
)

func main() {
	var configFilePath string
	flag.StringVar(&configFilePath, "config", "config.yml", "a path to the config file")
	flag.Parse()
	_, err := NewConfig(configFilePath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(configFilePath)
}
