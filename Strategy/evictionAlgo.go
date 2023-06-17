package Strategy

type EvictionAlgo interface {
	evict(c *Cache)
}
