package router

import (
	"github.com/GforsZi/gin-api/api/internal/handler"
	"github.com/gin-gonic/gin"
)

func New(userHandler *handler.UserHandler) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		users := api.Group("/users")
		{
			users.POST("/register", userHandler.Register)
			users.GET("/:id", userHandler.GetByID)
			users.GET("", userHandler.GetAll)
		}
	}

	return r
}
