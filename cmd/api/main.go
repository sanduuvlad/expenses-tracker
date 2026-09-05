package main

import (
	"expense-tracker/internal/config"
	"expense-tracker/internal/database"
	"expense-tracker/internal/handler"
	"expense-tracker/internal/repository"
	"expense-tracker/internal/service"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
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

	repo := repository.NewRepository(db)

	userService := service.NewUserService(repo)

	handler := handler.NewHandler(userService)

	router := gin.Default()

	router.GET("/users", handler.GetAllUsers)

	address := fmt.Sprintf(":%d", cfg.Server.Port)

	if err := router.Run(address); err != nil {
		log.Fatal(err)
	}
}
