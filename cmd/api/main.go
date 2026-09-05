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

	pool, err := database.Connect(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	// Repository
	userRepo := repository.NewUserRepository(pool)

	// Service
	userService := service.NewUserService(userRepo)

	// Handler
	userHandler := handler.NewUserHandler(userService)

	// Router
	router := gin.Default()

	router.GET("/users", userHandler.GetAllUsers)

	// Server
	address := fmt.Sprintf(":%d", cfg.Server.Port)

	if err := router.Run(address); err != nil {
		log.Fatal(err)
	}
}
