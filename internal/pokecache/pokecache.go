package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct{
	createdAt time.Time
	val []byte
}

type Cache struct {
	entries map[string]cacheEntry
	cacheLock sync.Mutex
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache{
	entry := map[string]cacheEntry{}
	c := &Cache{
		entries: entry,
		cacheLock: sync.Mutex{},
		interval: interval,
	}

	go c.reapLoop()

	return c
}

func (c *Cache) Add(key string, val []byte){
	c.cacheLock.Lock()
	defer c.cacheLock.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}

}

func (c *Cache) Get(key string) ([]byte,bool){
	c.cacheLock.Lock()
	entry := c.entries[key].val
	c.cacheLock.Unlock()
	if(entry!=nil){
		return entry,true
	}
	return nil,false
}

func (c *Cache) reapLoop(){
	for{
		time.Sleep(c.interval)

		c.cacheLock.Lock()
		for key,entry :=range c.entries{
			if(entry.createdAt.Add(c.interval).Before(time.Now())){
				delete(c.entries, key)
			}
		}
		c.cacheLock.Unlock()
	}
	
}