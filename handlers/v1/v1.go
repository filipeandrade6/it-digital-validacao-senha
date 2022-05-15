package v1

import (
	"net/http"

	"github.com/filipeandrade6/iti-digital-desafio-backend/data"
	"github.com/filipeandrade6/iti-digital-desafio-backend/validation"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// Validator is a struct that hold the ValidateHandle function with the validators declared
// in the NewValidator construction.
type Validator struct {
	log       *zap.Logger
	validator func(string) bool
}

// NewValidator creates a Validator with the provided ValidatorFuncs if provided, otherwise
// return the Defaultvalidator.
func NewValidator(log *zap.Logger, v ...func(func(string) bool) func(string) bool) *Validator {
	// Return the default validator if no validators was provided.
	if len(v) == 0 {
		return &Validator{log, validation.NewDefaultValidator()}
	}

	return &Validator{log, validation.NewValidator(v...)}
}

// ValidateHandle is a handler that verifies if the password provided in the request meets
// the criteria declared in the Validator construction.
func (v *Validator) ValidateHandle(c *gin.Context) {
	var jsonPass data.Password
	if err := c.ShouldBindJSON(&jsonPass); err != nil {
		v.log.Error(
			"invalid payload",
			zap.String("method", "POST"),
			zap.Int("status", http.StatusBadRequest),
			zap.String("ip", c.ClientIP()),
			zap.String("path", c.FullPath()),
		)
		c.JSON(http.StatusBadRequest, gin.H{"is_valid": false})
		return
	}

	if !v.validator(jsonPass.Password) {
		v.log.Error(
			"invalid password",
			zap.String("method", "POST"),
			zap.Int("status", http.StatusBadRequest),
			zap.String("ip", c.ClientIP()),
			zap.String("path", c.FullPath()),
		)
		c.JSON(http.StatusBadRequest, gin.H{"is_valid": false})
		return
	}

	v.log.Info(
		"successful request",
		zap.String("method", "POST"),
		zap.Int("status", http.StatusOK),
		zap.String("ip", c.ClientIP()),
		zap.String("path", c.FullPath()),
	)
	c.JSON(http.StatusOK, gin.H{"is_valid": true})
}
