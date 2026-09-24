package pokecache

import (
	"sync"
	"time"
)

// Structs
type Cache struct {
	cache map[string]cacheEntry
	mu     *sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// Functions
func NewCache(interval time.Duration) Cache {
	cache := Cache{
		cache: make(map[string]cacheEntry),
		mu:     &sync.RWMutex{},
	}

	go func() {
		ticker := time.NewTicker(interval)
		for range ticker.C {
			cache.reapLoop(interval)
		}
	}()

	return cache
}

func (C *Cache) Add(key string, val []byte) {
	C.mu.Lock()
	defer C.mu.Unlock()

	C.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (C *Cache) Get(key string) ([]byte, bool) {
	C.mu.RLock()
	defer C.mu.RUnlock()

	val, ok := C.cache[key]
	if !ok {
		return nil, ok
	}
	return val.val, ok
}

func (C *Cache) reapLoop(interval time.Duration) {
	C.mu.Lock()
	defer C.mu.Unlock()

	for k, v := range C.cache {
		timeNow := time.Now()
		timeSince := timeNow.Sub(v.createdAt)
		if timeSince >= interval {
			// Delete value
			delete(C.cache, k)
		}
	}
}
