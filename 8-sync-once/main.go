package main

import (
	"fmt"
	"sync"
)

var (
	config     map[string]string
	initConfig sync.Once
)

func loadConfig() {
	initConfig.Do(func() {
		fmt.Println("Init config...")
		config = map[string]string{
			"env":     "prod",
			"version": "1.0.0",
		}
	})
}

func main() {
	loadConfig() // executed once
	loadConfig() // ingored
	fmt.Printf("Config: %v\n", config)
}
