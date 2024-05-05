package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/phongtran11/go-project/handlers"
)

func AuthRoutes(r *gin.RouterGroup) {
	authHandlers := handlers.AuthHandler(r)

	r.POST("/login", authHandlers.Login)
	r.POST("/register", authHandlers.Register)
}
