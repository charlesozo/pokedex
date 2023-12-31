package cache

import (
	"time"
	// "sync"
)

type Cache struct{
	cache map[string]cacheEntry
	// mux   *sync.Mutex
}
type cacheEntry struct {
	createdAt time.Time
	val []byte
}
func NewCache(time time.Duration) Cache{
	c := Cache{
		cache: make(map[string]cacheEntry),
	}
	go c.reapLoop(time)
	return c
	// mux: &sync.Mutex{},
	
} 

func (c *Cache) Add(key string, val []byte){
 c.cache[key] = cacheEntry{
    val: val,
	createdAt: time.Now().UTC(),
 }
}
func (c *Cache) Get(key string) ([]byte, bool){
	cacheE, ok := c.cache[key]
	return cacheE.val, ok
}

func (c *Cache) reapLoop(interval time.Duration){
	ticker := time.NewTicker(interval)
	for range ticker.C{
      c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration){
	timeAgo := time.Now().UTC().Add(-interval)
	for k, v := range c.cache{
		if v.createdAt.Before(timeAgo){
			delete(c.cache, k)
		}
	}
}