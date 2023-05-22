package memorycache

import "sync"

// Cache struct of memory cache object
type Cache struct {
	sync.RWMutex
	items map[string]Item
}

// Item struct for value of cache
type Item struct {
	Value interface{}
}

func New() *Cache {

	items := make(map[string]Item)

	cache := Cache{
		items: items,
	}

	return &cache
}

// Set set a key and value in cache
func (c *Cache) Set(key string, value interface{}) {

	c.Lock()

	defer c.Unlock()

	c.items[key] = Item{
		Value: value,
	}
}

// Get get value from cache by key
func (c *Cache) Get(key string) (interface{}, bool) {

	c.RLock()

	defer c.RUnlock()

	item, found := c.items[key]

	// ключ не найден
	if !found {
		return nil, false
	}

	return item.Value, true
}
