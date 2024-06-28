package main

import "fmt"

// Step 1 : Strategy defines the interface for different strategies.
type Strategy interface {
	Execute(a, b int) int
}

// 2. Define concrete strategies

// Addition Strategy
type Addition struct{}

func (a Addition) Execute(x, y int) int {
	return x + y
}

// Subtraction strategy
type Subtraction struct{}

func (s Subtraction) Execute(x, y int) int {
	return x - y
}

// Step 3 : Define a context that maintains a reference to a Strategy object.

// Context maintains a reference to a strategy object.
type Context struct {
	strategy Strategy
}

// SetStrategy allows changing the strategy at runtime.
func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

// ExecuteStrategy executes the strategy.
func (c *Context) ExecuteStrategy(a, b int) int {
	return c.strategy.Execute(a, b)
}

func main() {

	// Step 4 : Using the strategy
	context := &Context{}

	// Using the addition strategy
	context.SetStrategy(Addition{})

	// Meaning in which context you're executing the strategy
	fmt.Println("Addition:", context.ExecuteStrategy(5, 3)) // Output: Addition: 8

	// Using the subtraction strategy
	context.SetStrategy(Subtraction{})
	fmt.Println("Subtraction:", context.ExecuteStrategy(5, 3)) // Output: Subtraction: 2

}
