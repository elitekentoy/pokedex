package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	data  map[string]CacheEntry
	mutex *sync.Mutex
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		data:  make(map[string]CacheEntry),
		mutex: &sync.Mutex{},
	}

	go cache.reapLoop(interval)

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	c.data[key] = CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mutex.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	entry, found := c.data[key]
	c.mutex.Unlock()
	if found {
		return entry.val, found
	}

	return []byte{}, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for key, value := range c.data {
		if value.createdAt.Before(now.Add(-last)) {
			delete(c.data, key)
		}
	}
}
