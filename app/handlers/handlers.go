package handlers

import (
	v1 "github.com/filipeandrade6/iti-digital-desafio-backend/app/handlers/v1"

	"github.com/gin-gonic/gin"
)

// NewAPI create a router with the routes and middlewares attached.
func NewAPI() *gin.Engine {
	router := gin.Default()

	// v1 routes.
	v1grp := router.Group("/v1")
	{
		vld := v1.NewValidator()
		v1grp.POST("/validate", vld.ValidateHandle)
	}

	return router
}
