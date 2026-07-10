package main

import (
	"context"
	"log"

	"github.com/GforsZi/gin-api/api/internal/config"
	"github.com/GforsZi/gin-api/api/internal/database"
	"github.com/GforsZi/gin-api/api/internal/firebase"
	"github.com/GforsZi/gin-api/api/internal/handler"
	"github.com/GforsZi/gin-api/api/internal/model"
	"github.com/GforsZi/gin-api/api/internal/repository"
	"github.com/GforsZi/gin-api/api/internal/router"
	"github.com/GforsZi/gin-api/api/internal/service"
)

func main() {
	cfg := config.Load()
	db := database.Connect(cfg)

	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	firebaseAuth, err := firebase.InitAuth(context.Background(), cfg.FirebaseCredentialsPath)
	if err != nil {
		log.Fatalf("firebase init failed: %v", err)
	}

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService, firebaseAuth)

	r := router.New(userHandler, firebaseAuth)

	log.Printf("Server running on port %s", cfg.AppPort)
	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Fatal(err)
	}

}
