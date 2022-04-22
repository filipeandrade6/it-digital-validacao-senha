package v1

import (
	"net/http"

	"github.com/filipeandrade6/iti-digital-desafio-backend/data"
	"github.com/filipeandrade6/iti-digital-desafio-backend/validation"

	"github.com/gin-gonic/gin"
)

// Validator is a struct that hold the ValidateHandle function with the validators declared
// in the NewValidator construction.
type Validator struct {
	validator func(string) bool
}

// NewValidator creates a Validator with the provided ValidatorFuncs if provided, otherwise
// return the Defaultvalidator.
func NewValidator(vls ...func(func(string) bool) func(string) bool) *Validator {
	// Return the default validator if no validators was provided.
	if len(vls) == 0 {
		return &Validator{validation.NewDefaultValidator()}
	}

	return &Validator{validation.NewValidator(vls...)}
}

// ValidateHandle is a handler that verifies if the password provided in the request meets
// the criteria declared in the Validator construction.
func (v *Validator) ValidateHandle(c *gin.Context) {
	var jsonPass data.Password
	if err := c.ShouldBindJSON(&jsonPass); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "invalid payload"})
		return
	}

	if !v.validator(jsonPass.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"status": "invalid password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "good to go :)"})
}
