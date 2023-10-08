package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	data map[string]CacheEntry
	mu   sync.RWMutex
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) add(key string, val []byte) {
	newEntry := CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	c.data[key] = newEntry
}

func (c *Cache) get(key string) ([]byte, bool) {
	v, exists := c.data[key]
	if !exists {
		return []byte{}, false
	}
	return v.val, true
}
