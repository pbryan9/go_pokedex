package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type PokeCache struct {
	Cache map[string]cacheEntry
	mux   sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *PokeCache {
	cache := PokeCache{
		Cache: make(map[string]cacheEntry, 0),
		mux:   sync.RWMutex{},
	}
	go cache.reapLoop(interval)
	return &cache
}

func (c *PokeCache) Add(key string, val []byte) {
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mux.Lock()
	c.Cache[key] = entry
	c.mux.Unlock()
	fmt.Printf("Added entry %s to cache\n", key)
}

func (c *PokeCache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		<-ticker.C
		c.mux.Lock()
		for k, e := range c.Cache {
			if time.Since(e.createdAt) > interval {
				fmt.Printf("deleting %s from cache\n", k)
				delete(c.Cache, k)
			}
		}
		c.mux.Unlock()
	}
}

func (c *PokeCache) Get(key string) ([]byte, bool) {
	c.mux.RLock()
	val, ok := c.Cache[key]
	c.mux.RUnlock()
	if !ok {
		return nil, false
	}

	return val.val, true
}
