package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(engine *gin.Engine) {
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	})
}