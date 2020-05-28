package config

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine()  {
	router := gin.Default()
	router.GET("/", welcome)
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To API",
	})
	return
}
