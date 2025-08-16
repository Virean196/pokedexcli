package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache    map[string]CacheEntry
	mu       sync.Mutex
	interval time.Duration
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(inter time.Duration) *Cache {
	cache := &Cache{
		cache:    make(map[string]CacheEntry),
		mu:       sync.Mutex{},
		interval: inter,
	}
	go cache.reapLoop()
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.cache[key] = CacheEntry{time.Now(), val}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	entry, exists := c.cache[key]
	defer c.mu.Unlock()
	if exists {
		return entry.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.mu.Lock()
		for key := range c.cache {
			if time.Since(c.cache[key].createdAt) > c.interval {
				delete(c.cache, key)
			}
		}
		c.mu.Unlock()
	}
}
