package main

// Evacuation Algorithm
type EvictionAlgo interface {
	evict(c *Cache)
}
