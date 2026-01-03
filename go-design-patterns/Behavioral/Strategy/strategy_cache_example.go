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
type EvictStrategy interface {
	Evict() // it's your choice to keep pass *Cache in Evict method
}

// Concrete Strategy
type LRUStrategy struct{}

func (l LRUStrategy) Evict() {
	println("Evicting by LRU strategy")
}

// Concrete strategy
type LFUStrategy struct{}

func (l LFUStrategy) Evict() {
	println("Evicting by LFU strategy")
}

// Context :: CACHE
// Cache is using Eviction Strategy to evict items from the cache
// Notice it holds the interface, not a concrete implementation
type Cache struct {
	strategy EvictStrategy
}

func NewCache(s EvictStrategy) *Cache {
	return &Cache{strategy: s}
}

func (c *Cache) Do() {
	c.strategy.Evict()
}

func main() {

	// Using the LRU strategy
	cache := NewCache(LRUStrategy{})
	cache.Do() // Output: Evicting by LRU strategy

	// Using the LFU strategy
	cache = NewCache(LFUStrategy{})
	cache.Do() // Output: Evicting by LFU strategy

}
