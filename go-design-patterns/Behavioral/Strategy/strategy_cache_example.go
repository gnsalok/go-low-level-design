package main

/*
What you need :-
1. Context :: Cache
2. Strategy :: EvictAlgo
3. Concrete Strategies :: LRU and FIFO
4. Context (Cache) will delegate the execution of strategy

Why use it?
- Strategy pattern allows you to define a family of algorithms, encapsulate each one, and make them interchangeable.
- Strategy lets the algorithm vary independently from clients that use it.
- In this example, we can easily add new eviction strategies (like LFU) without modifying the Cache class.
- The client code (main function) can switch between different eviction strategies at runtime.

*/

// Strategy Interface
type EvictAlgo interface {
	Evict(c *Cache) // it's your choice to keep pass *Cache in Evict method
}

// Concrete Strategy
type LRU struct{}

func (l LRU) Evict(c *Cache) {
	println("Evicting by LRU strategy")
}

// Concrete strategy
type LFU struct{}

func (l LFU) Evict(c *Cache) {
	println("Evicting by LFU strategy")
}

// Context :: CACHE

type Cache struct {
	evictAlgo EvictAlgo
}

func (c *Cache) setEvictionAlgo(e EvictAlgo) {
	c.evictAlgo = e
}

func main() {

	// Using the LRU strategy
	cache := &Cache{}
	cache.setEvictionAlgo(LRU{})
	cache.evictAlgo.Evict(cache) // Output: Evicting by LRU strategy

	// Using the LFU strategy
	cache.setEvictionAlgo(LFU{})
	cache.evictAlgo.Evict(cache) // Output: Evicting by LFU strategy

}
