package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phongtran11/go-project/pkg/constants"
	"github.com/phongtran11/go-project/pkg/dto/request"
	"github.com/phongtran11/go-project/pkg/helpers"
	"github.com/phongtran11/go-project/services"
)

type TAuthHandler struct {
	AuthServices *services.TAuthServices
}

func AuthHandler(r *gin.RouterGroup) *TAuthHandler {

	return &TAuthHandler{
		AuthServices: services.AuthServices(),
	}
}

func (authHandler *TAuthHandler) Login(c *gin.Context) {
	request := request.TLoginRequest{}

	// Validate request
	validationError := c.ShouldBindJSON(&request)

	if validationError != nil {
		c.JSON(http.StatusBadRequest, helpers.GenerateBaseResponseWithValidationError(nil, false, constants.ValidationError, validationError))
		return
	}

	// process login
	tokenResponse, servicesError := authHandler.AuthServices.Login(request)

	// handler service error
	if servicesError != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.GenerateBaseResponseWithError(nil, false, constants.ValidationError, servicesError))
		return
	}

	c.JSON(http.StatusOK, helpers.GenerateBaseResponse(tokenResponse, true))
}

func (authHandler *TAuthHandler) Register(c *gin.Context) {
	request := request.TRegisterRequest{}

	// Validate request
	validationError := c.ShouldBindJSON(&request)

	if validationError != nil {
		c.JSON(http.StatusBadRequest, helpers.GenerateBaseResponseWithValidationError(nil, false, constants.ValidationError, validationError))
		return
	}

	// process register
	servicesError := authHandler.AuthServices.Register(request)

	// handler service error
	if servicesError != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.GenerateBaseResponseWithError(nil, false, constants.ValidationError, servicesError))
		return
	}

	c.JSON(http.StatusOK, helpers.GenerateBaseResponse(nil, true))
}
