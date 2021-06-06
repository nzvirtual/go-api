package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/common-nighthawk/go-figure"
	log "github.com/dhawton/log4g"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nzvirtual/go-api/lib/cache"
	"github.com/nzvirtual/go-api/lib/database/models"
)

func main() {
	log.SetLogLevel(log.DEBUG)

	intro := figure.NewFigure("NZV API", "", false).Slicify()
	for i := 0; i < len(intro); i++ {
		log.Category("main").Info(intro[i])
	}

	log.Category("main").Info("Starting NZV API")
	log.Category("main").Info("Checking for .env, loading if exists")
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load()
		if err != nil {
			log.Category("main").Fatal("Error loading .env file " + err.Error())
		}
	}

	appenv := Getenv("APP_ENV", "dev")

	log.Category("main").Info(fmt.Sprintf("APP_ENV=%s", appenv))

	if appenv == "production" {
		log.SetLogLevel(log.INFO)
		log.Category("main").Info("Setting gin to Release Mode")
		gin.SetMode(gin.ReleaseMode)
	} else {
		log.SetLogLevel(log.DEBUG)
	}

	log.Category("main").Info("Loading cache server")
	port, err := strconv.Atoi(Getenv("CACHE_PORT", "0"))
	if err != nil {
		log.Category("main").Error("Invalid port specified, could not convert to int")
		port = 0
	}

	cache.Connect(&cache.Options{
		Driver:   Getenv("CACHE_DRIVER", ""),
		Hostname: Getenv("CACHE_HOSTNAME", "localhost"),
		Port:     port,
		Username: Getenv("CACHE_USERNAME", ""),
		Password: Getenv("CACHE_PASSWORD", ""),
	})

	log.Category("main").Info("Configuring Gin Server")
	server := NewServer(appenv)

	log.Category("main").Info("Connecting to database and handling migrations")
	models.Connect(Getenv("DB_USERNAME", "root"), Getenv("DB_PASSWORD", "secret"), Getenv("DB_HOSTNAME", "localhost"), Getenv("DB_PORT", "3306"), Getenv("DB_DATABASE", "nzvirtual"))
	models.DB.AutoMigrate()

	log.Category("main").Info("Done with setup, starting web server...")
	server.engine.Run(fmt.Sprintf(":%s", Getenv("PORT", "3000")))
}

func Getenv(key string, defaultValue string) string {
	val := os.Getenv(key)
	if len(val) == 0 {
		return defaultValue
	}
	return val
}
