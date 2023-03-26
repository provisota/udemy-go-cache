package cache

type Cache struct {
	cacheMap map[string]any
}

func (c *Cache) Set(key string, value any) {
	c.cacheMap[key] = value
}

func (c *Cache) Get(key string) any {
	return c.cacheMap[key]
}

func (c *Cache) Delete(key string) {
	delete(c.cacheMap, key)
}

func New() *Cache {
	return &Cache{make(map[string]any)}
}
