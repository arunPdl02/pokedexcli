package pokecache

import "time"

func NewCache(interval time.Duration) *Cache {
	my_cache := &Cache{
		cache:    make(map[string]CacheEntry),
		interval: interval,
	}
	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			my_cache.reapLoop()
		}
	}()
	return my_cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.cache[key]
	if !ok {
		return nil, false
	}
	return val.val, true
}

func (c *Cache) reapLoop() {
	var key_to_del []string

	c.mu.Lock()
	defer c.mu.Unlock()

	for key, value := range c.cache {
		if time.Since(value.createdAt) > c.interval {
			key_to_del = append(key_to_del, key)
		}
	}
	for _, key := range key_to_del {
		delete(c.cache, key)
	}
}
