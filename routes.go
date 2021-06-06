package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "github.com/nzvirtual/go-api/controllers/v1"
)

func SetupRoutes(engine *gin.Engine) {
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "PONG"})
	})

	v1Router := engine.Group("/v1")
	{
		airportRouter := v1Router.Group("/airport")
		{
			airportRouter.GET("/:icao", v1.GetAirport)
		}

		authRouter := v1Router.Group("/auth")
		{
			authRouter.POST("/register", v1.PostRegister)
			authRouter.POST("/login", v1.PostLogin)
		}
	}
}
