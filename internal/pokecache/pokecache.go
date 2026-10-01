package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache    map[string]cacheEntry
	mu       *sync.Mutex
	interval time.Duration
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry := cacheEntry{
		val:       val,
		createdAt: time.Now(),
	}

	c.cache[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	found, exists := c.cache[key]
	return found.val, exists
}

func (c *Cache) reapLoop() {

	ticker := time.NewTicker(c.interval)

	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.cache {

			subSeconds := time.Until(entry.createdAt) * -1

			if subSeconds >= c.interval {
				delete(c.cache, key)
			}
		}
		c.mu.Unlock()
	}
}

func NewCache(interval time.Duration) *Cache {
	c := Cache{
		cache:    make(map[string]cacheEntry),
		mu:       &sync.Mutex{},
		interval: interval,
	}

	go c.reapLoop()

	return &c

}
