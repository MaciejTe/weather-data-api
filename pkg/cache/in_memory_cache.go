package cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

// InMemoryCache implements go-cache implementation of cache interface.
type InMemoryCache struct {
	cache *cache.Cache
}

// NewCache creates in-memory cache object and returns it.
func NewCache(defaultExpiration, cleanupInterval time.Duration) Cache {
	inMemoryCache := InMemoryCache{
		cache: cache.New(defaultExpiration, cleanupInterval),
	}
	return &inMemoryCache
}

// Get retrieves data from in-memory cache.
func (i *InMemoryCache) Get(key string) (interface{}, bool) {
	return i.cache.Get(key)
}

// Set saves data to in-memory cache with default go-cache expiration time.
func (i *InMemoryCache) Set(key string, value interface{}) {
	i.cache.SetDefault(key, value)
}
