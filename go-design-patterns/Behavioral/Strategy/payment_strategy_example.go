package main

import "fmt"

/*
Strategy Pattern:
- Strategy pattern allows you to define a family of algorithms, encapsulate each one, and make them interchangeable.
- What? - used for interchangeable behaviors
- Used for - Pricing, Auth, Payments etc.
*/

// PayStrategy defines the behavior (strategy interface)
type PayStrategy interface {
	Pay(amount int) error
}

// Stripe is one concrete type
type Stripe struct{}

func (Stripe) Pay(amount int) error {
	fmt.Printf("Paid %d using Stripe\n", amount)
	return nil
}

// PayPal is another concrete strategy
type PayPal struct{}

func (PayPal) Pay(amount int) error {
	fmt.Printf("Paid %d using PayPal\n", amount)
	return nil
}

// Checkout is the context that uses a strategy
type Checkout struct {
	strategy PayStrategy
}

// NewCheckout sets the strategy
func NewCheckout(s PayStrategy) *Checkout {
	return &Checkout{strategy: s}
}

// Do executes the payment via chosen strategy
func (c *Checkout) Do(amount int) error {
	return c.strategy.Pay(amount)

}

func main() {
	stripeCheckout := NewCheckout(Stripe{})
	stripeCheckout.Do(100)

	// Swap to PayPal strategy
	paypalCheckout := NewCheckout(PayPal{})
	paypalCheckout.Do(200)
}
