package pokecache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) *Cache {
	my_cache := &Cache{
		cache: make(map[string]CacheEntry),
		ttl:   interval,
		mu:    &sync.Mutex{},
	}
	go my_cache.reapLoop(interval)
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
	return val.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap()
	}
}

func (c *Cache) reap() {
	var key_to_del []string

	c.mu.Lock()
	defer c.mu.Unlock()

	for key, value := range c.cache {
		if time.Since(value.createdAt) > c.ttl {
			key_to_del = append(key_to_del, key)
		}
	}
	for _, key := range key_to_del {
		delete(c.cache, key)
	}
}
