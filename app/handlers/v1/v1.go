package v1

import (
	"net/http"

	"github.com/filipeandrade6/iti-digital-desafio-backend/data"
	"github.com/filipeandrade6/iti-digital-desafio-backend/validation"

	"github.com/gin-gonic/gin"
)

// * Nomear com handler...

type Validator struct{}

func NewValidator() *Validator {
	return &Validator{}
}

func (v *Validator) Validate(c *gin.Context) {
	var jsonPass data.Password
	if err := c.ShouldBindJSON(&jsonPass); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !validation.IsValid(jsonPass.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "your pass should have..."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}
