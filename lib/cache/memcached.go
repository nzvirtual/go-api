package cache

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

type MemcacheCache struct {
	mdb *memcache.Client
}

func ConnectMemcache(options *Options) (*MemcacheCache, error) {
	mc := &MemcacheCache{}
	mc.mdb = memcache.New(fmt.Sprintf("%s:%d", options.Hostname, options.Port))

	return mc, nil
}

func (m *MemcacheCache) Set(key string, value string) (bool, error) {
	m.mdb.Set(&memcache.Item{Key: key, Value: []byte(value)})

	return true, nil
}

func (m *MemcacheCache) Get(key string, defaultValue string) (string, error) {
	it, err := m.mdb.Get(key)

	if err != nil {
		return "", err
	}

	return string(it.Value), nil
}
