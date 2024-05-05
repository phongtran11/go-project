package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phongtran11/go-project/pkg/dto/response"
	"github.com/phongtran11/go-project/pkg/helpers"
)

type THealthHandler struct{}

func HealthHandler() *THealthHandler {
	return &THealthHandler{}
}

func (healthHandler *THealthHandler) Health(c *gin.Context) {
	healthResponse := response.THealthResponse{
		Message: "OK",
	}

	c.JSON(http.StatusCreated, helpers.GenerateBaseResponse(healthResponse, true))
}
