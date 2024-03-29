package main

type Cache struct {
	storage       map[string]string
	evictiongAlgo EvictionAlgo
	capacity      int
	maxCapacity   int
}

func initCache(e EvictionAlgo) *Cache {
	storage := make(map[string]string)
	return &Cache{
		storage:       storage,
		evictiongAlgo: e,
		capacity:      0,
		maxCapacity:   2,
	}
}

func (c *Cache) setEvacuationAlgo(e EvictionAlgo) {
	c.evictiongAlgo = e
}

func (c *Cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *Cache) get(key string) {
	delete(c.storage, key)
}

func (c *Cache) evict() {
	c.evictiongAlgo.evict(c)
	c.capacity--
}
