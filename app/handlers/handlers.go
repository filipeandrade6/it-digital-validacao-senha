package handlers

import (
	v1 "github.com/filipeandrade6/iti-digital-desafio-backend/app/handlers/v1"

	"github.com/gin-gonic/gin"
)

func NewAPI() *gin.Engine {
	// Default with the Logger and Recovery middleware already attached.
	router := gin.Default()

	// v1 routes.
	v1grp := router.Group("/v1")
	{
		vld := v1.NewValidator()
		v1grp.POST("/validate", vld.Validate)
	}

	return router
}
