package cache

// Cache interface helps to represent any kind of Caching implementation.
type Cache interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{})
}
