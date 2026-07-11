package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]CacheEntry
	ttl   time.Duration
	mu    *sync.Mutex
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}
