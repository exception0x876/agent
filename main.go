package main

import (
	"flag"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	var configFilePath string
	flag.StringVar(&configFilePath, "config", "config.yml", "a path to the config file")
	flag.Parse()
	config, err := NewConfig(configFilePath)
	if err != nil {
		fmt.Println(err)
	}
	wg.Add(1)
	go monitorSmart(config)
	wg.Wait()
	fmt.Println("Main thread finished")
}
