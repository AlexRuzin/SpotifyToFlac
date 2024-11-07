package main

import (
	"fmt"
)

func main() {
	path := "config.ini"
	config, err := LoadConfig(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Successfully loaded file: " + config.filePath)

	return
}
