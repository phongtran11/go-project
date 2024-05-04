package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phongtran11/go-project/dto/response"
	"github.com/phongtran11/go-project/helpers"
)

func HealthHandler(r *gin.RouterGroup) {

	r.GET("/", func(c *gin.Context) {
		healthResponse := response.THealthResponse{
			Message: "OK",
		}

		c.JSON(http.StatusCreated, helpers.GenerateBaseResponse(healthResponse, true))
	})
}
