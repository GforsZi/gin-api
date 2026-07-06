package main

import (
	"log"

	"github.com/GforsZi/gin-api/api/internal/config"
	"github.com/GforsZi/gin-api/api/internal/database"
	"github.com/GforsZi/gin-api/api/internal/handler"
	"github.com/GforsZi/gin-api/api/internal/model"
	"github.com/GforsZi/gin-api/api/internal/repository"
	"github.com/GforsZi/gin-api/api/internal/router"
	"github.com/GforsZi/gin-api/api/internal/service"
)

// type album struct {
// 	ID     string  `json:"id"`
// 	Title  string  `json:"title"`
// 	Artist string  `json:"artist"`
// 	Price  float64 `json:"price"`
// }
//
// var albums = []album{
// 	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
// 	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
// 	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
// }
//
// func getAlbums(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, albums)
// }

func main() {
	cfg := config.Load()
	db := database.Connect(cfg)

	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r := router.New(userHandler)

	log.Printf("Server running on port %s", cfg.AppPort)
	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Fatal(err)
	}

}
