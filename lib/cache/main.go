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
	var c ICache
	var err error

	if strings.ToLower(options.Driver) == "redis" {
		log.Category("cache").Debug("Configuring redis driver")
		c, err = ConnectRedis(options)
		if err != nil {
			log.Category("cache").Fatal("Error connecting to Redis " + err.Error())
			panic("Error connecting to redis")
		}
	} else if strings.ToLower(options.Driver) == "memcache" {
		log.Category("cache").Debug("Configuring memcache driver")
		c, err = ConnectMemcache(options)
		if err != nil {
			log.Category("cache").Fatal("Error configuring Memcache " + err.Error())
			panic("Error connecting to memcache")
		}
	} else {
		log.Category("cache").Fatal("No/Invalid cache driver configured")
		panic("No/Invalid cache driver configured")
	}

	cache = c
}

func Get(key string, defaultValue string) (string, error) {
	return cache.Get(key, defaultValue)
}

func Set(key string, value string) (bool, error) {
	return cache.Set(key, value)
}
