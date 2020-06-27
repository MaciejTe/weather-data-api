package cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

// InMemoryCache implements go-cache implementation of cache interface.
type InMemoryCache struct {
	cache *cache.Cache
}

func NewCache(defaultExpiration, cleanupInterval time.Duration) Cache {
	inMemoryCache := InMemoryCache {
		cache: cache.New(defaultExpiration, cleanupInterval),
	}
	return &inMemoryCache
}

func (i *InMemoryCache) Get(key string) (interface{}, bool) {
	return i.cache.Get(key)
}

func (i *InMemoryCache) Set(key string, value interface{}) {
	i.cache.SetDefault(key, value)
}
