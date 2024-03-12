package cacheLib

import (
	"errors"
	"sync"
	"time"
)

type Cache struct {
	cache map[string]Profile
	mu    sync.RWMutex
	ttl   time.Duration
}

type Profile struct {
	Data    string
	SysInfo SysInfo
}

type SysInfo struct {
	CreatedAt time.Time
}

func NewCache(ttl time.Duration) *Cache {
	cache := &Cache{
		cache: make(map[string]Profile),
		mu:    sync.RWMutex{},
		ttl:   ttl,
	}
	go cache.StartCleanUp()
	return cache
}

func (c *Cache) StartCleanUp() {
	ticker := time.NewTicker(c.ttl)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			c.Clean()
		}
	}
}

func (c *Cache) Set(key string, value Profile) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = value
}

func (c *Cache) Get(key string) interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, ok := c.cache[key]
	if !ok {
		return errors.New("Key didn't find")
	}
	return value
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.cache, key)
}

func (c *Cache) Clean() {
	c.mu.Lock()
	defer c.mu.Unlock()
	now := time.Now()
	for id, profile := range c.cache {
		if now.Sub(profile.SysInfo.CreatedAt) >= c.ttl {
			delete(c.cache, id)
		}
	}
}
