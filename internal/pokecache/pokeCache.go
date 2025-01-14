package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache    map[string]cacheEntry
	mu       *sync.Mutex
	interval time.Duration
}

func (c Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{
		time.Now(),
		value,
	}
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, exists := c.cache[key]
	if !exists {
		return nil, false
	}
	return entry.val, true
}

func (c Cache) reapLoop(){
	clock := time.NewTicker(c.interval)
	defer clock.Stop()
	for range clock.C{
		c.mu.Lock()
		for k,v:= range c.cache{
			if time.Since(v.createdAt)>=c.interval{
				delete (c.cache,k)
			}
		}
		c.mu.Unlock()
	}
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval int) Cache {
	return Cache{
		make(map[string]cacheEntry),
		&sync.Mutex{},
		time.Duration(interval) * time.Second,
	}
}
