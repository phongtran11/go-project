package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/phongtran11/go-project/handlers"
)

func InitRoutes() {
	gin.SetMode(gin.DebugMode)

	r := gin.Default()

	initRoutesWithEngine(r)

	err := r.Run(":3000")

	if err != nil {
		panic("Failed to start server")
	}
}

func initRoutesWithEngine(engine *gin.Engine) {
	// define base route
	api := engine.Group("/api")

	// define version 1 route
	v1 := api.Group("/v1")

	health := v1.Group("/health")
	handlers.HealthHandler(health)

	auth := v1.Group("/auth")
	handlers.AuthHandler(auth)
}
