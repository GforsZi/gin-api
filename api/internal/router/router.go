package router

import (
	"github.com/GforsZi/gin-api/api/internal/handler"
	"github.com/GforsZi/gin-api/api/internal/middleware"
	firebaseAuth "firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
)

func New(userHandler *handler.UserHandler, authClient *firebaseAuth.Client) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		api.POST("/auth/firebase", userHandler.AuthFirebase)

		auth := api.Group("/auth")
		auth.Use(middleware.AuthMiddleware(authClient))
		{
			auth.GET("/me", userHandler.GetMe)
		}

		users := api.Group("/users")
		{
			users.POST("/register", userHandler.Register)
			users.GET("/:id", userHandler.GetByID)
			users.GET("", userHandler.GetAll)
		}
	}

	return r
}
