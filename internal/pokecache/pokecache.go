package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

type Cache struct {
	entry map[string]cacheEntry
	mu   sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	newCache := &Cache{
		entry: make(map[string]cacheEntry),
		mu: sync.Mutex{},
	}
	go newCache.reapLoop(interval)
	return newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.entry[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	cacheEntry, ok := c.entry[key]
	c.mu.Unlock()
	return cacheEntry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		c.mu.Lock()
		for k, v := range c.entry {
				if time.Since(v.createdAt) > interval {
					delete(c.entry, k)
				}
			}
		c.mu.Unlock()
	}	
}

