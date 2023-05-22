package memorycache

import "sync"

type Cache struct {
	sync.RWMutex
	items map[string]Item
}

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

func (c *Cache) Set(key string, value interface{}) {

	c.Lock()

	defer c.Unlock()

	c.items[key] = Item{
		Value: value,
	}
}

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
