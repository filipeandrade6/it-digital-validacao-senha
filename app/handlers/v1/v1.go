package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const tipMessage string = "password must have..."

func Routes(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.POST("/validate", validate)
	}
}

type Pass struct {
	Password string `json:"password" binding:"required"`
}

func validate(c *gin.Context) {
	var jsonPass Pass
	if err := c.ShouldBindJSON(&jsonPass); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !isValid(jsonPass.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": tipMessage})
	}

	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}

func isValid(password string) bool {
	return true
}
