package cache

import (
	"strings"

	log "github.com/dhawton/log4g"
)

type Options struct {
	Driver   string
	Hostname string
	Port     int
	Username string
	Password string
}

type ICache interface {
	Get(string, string) (string, error)
	Set(string, string) (bool, error)
}

var cache ICache

func Connect(options *Options) *ICache {
	if strings.ToLower(options.Driver) == "redis" {
		log.Category("cache").Debug("Connecting to Redis")
		c, err := ConnectRedis(options)
		log.Category("cache").Debug("Connection ready")
		if err != nil {
			log.Category("cache").Fatal("Error connecting to Redis " + err.Error())
			panic("Error connecting to redis")
		}

		cache = c
	}
}

func Get(key string, defaultValue string) (string, error) {
	return cache.Get(key, defaultValue)
}

func Set(key string, value string) (bool, error) {
	return cache.Set(key, value)
}
