package main

import (
	"fmt"
	"jeanmrtns/sample-go-api/config"
	"jeanmrtns/sample-go-api/router"
)

func main() {
	err := config.InitConfigs()

	if err != nil {
		fmt.Printf("Error while setting up project %v", err)
	}

	router.Initialize()
}
