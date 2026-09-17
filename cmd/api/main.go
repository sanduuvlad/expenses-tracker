package main

import (
	"expense-tracker/internal/config"
	"expense-tracker/internal/database"
	"expense-tracker/internal/handler"
	"expense-tracker/internal/middleware"
	"expense-tracker/internal/repository"
	"expense-tracker/internal/service"
	"expense-tracker/internal/token"
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

	// JWT Token
	tokenManager := token.NewTokenManager(cfg.JWT.Secret)

	// Repository
	userRepo := repository.NewUserRepository(pool)
	categoryRepo := repository.NewCategoryRepository(pool)

	// Service
	userService := service.NewUserService(userRepo, tokenManager)
	categoryService := service.NewCategoryService(categoryRepo)

	// Handler
	userHandler := handler.NewUserHandler(userService)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	// Router
	router := gin.Default()

	router.GET("/users", userHandler.GetAllUsers)
	router.GET("/users/:id", userHandler.GetUserByID)
	router.POST("/users", userHandler.RegisterUser)
	router.POST("/login", userHandler.LoginUser)

	// Auth
	authorized := router.Group("/")
	authorized.Use(middleware.AuthMiddleware(tokenManager))

	// AuthRouter
	authorized.POST("/categories", categoryHandler.CreateCategory)
	authorized.GET("/categories", categoryHandler.GetCategories)
	authorized.PATCH("/categories/:id", categoryHandler.UpdateCategory)
	authorized.DELETE("/categories/:id", categoryHandler.DeleteCategory)

	// Server
	address := fmt.Sprintf(":%d", cfg.Server.Port)

	if err := router.Run(address); err != nil {
		log.Fatal(err)
	}
}
