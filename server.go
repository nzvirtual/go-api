package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nzvirtual/go-api/middleware"
)

type Server struct {
	engine *gin.Engine
}

func NewServer(appenv string) *Server {
	server := Server{}

	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(middleware.Logger)
	server.engine = engine

	SetupRoutes(engine)

	return &server
}
