// Package handlers manages the different versions and endpoints of the API.
package handlers

import (
	v1 "github.com/filipeandrade6/iti-digital-desafio-backend/handlers/v1"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// NewAPI create a router with the routes and middlewares attached.
func NewAPI(log *zap.Logger) *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())

	// v1 routes.
	v1grp := router.Group("/v1")
	{
		vld := v1.NewValidator(log)
		v1grp.POST("/validate", vld.ValidateHandle)
	}

	return router
}
