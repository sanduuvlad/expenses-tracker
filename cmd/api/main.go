package main

import (
	"expense-tracker/internal/config"
	"fmt"
	"log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg.Server.Port)
}
