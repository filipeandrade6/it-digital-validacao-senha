package v1

import (
	"net/http"

	"github.com/filipeandrade6/iti-digital-desafio-backend/business/validation"

	"github.com/gin-gonic/gin"
)

const tipMessage string = "password must have..."

type Pass struct {
	Password string `json:"password" binding:"required"`
}

func Validate(c *gin.Context) {
	var jsonPass Pass
	if err := c.ShouldBindJSON(&jsonPass); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !validation.IsValid(jsonPass.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": tipMessage})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}
