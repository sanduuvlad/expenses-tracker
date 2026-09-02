package main

import (
	"expense-tracker/internal/config"
	"expense-tracker/internal/database"
	"fmt"
	"log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
	)

	db, err := database.Open(dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
