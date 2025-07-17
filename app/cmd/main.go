package main

import (
	"fmt"
	"mp-service/internal/config"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(cfg)
}
