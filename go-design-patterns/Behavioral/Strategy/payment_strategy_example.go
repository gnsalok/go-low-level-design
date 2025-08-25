package main

import "fmt"

// 1. The Strategy Interface
// This defines what any payment algorithm must be able to do.
type PaymentStrategy interface {
	Pay(amount float64) string
}

// 2. Concrete Strategies
// Each struct implements the PaymentStrategy interface.

// CreditCardStrategy is a specific payment method.
type CreditCardStrategy struct {
	CardNumber string
	Owner      string
}

func (cc *CreditCardStrategy) Pay(amount float64) string {
	return fmt.Sprintf("Paying $%.2f using Credit Card %s", amount, cc.CardNumber)
}

// PayPalStrategy is another specific payment method.
type PayPalStrategy struct {
	Email string
}

func (pp *PayPalStrategy) Pay(amount float64) string {
	return fmt.Sprintf("Paying $%.2f using PayPal account %s", amount, pp.Email)
}

// 3. The Context
// ShoppingCart uses a PaymentStrategy to process the payment.
// Notice it holds the interface, not a concrete type.
type ShoppingCart struct {
	totalAmount     float64
	paymentStrategy PaymentStrategy
}

// SetPaymentStrategy allows changing the payment method at runtime.
func (sc *ShoppingCart) SetPaymentStrategy(strategy PaymentStrategy) {
	sc.paymentStrategy = strategy
}

// Checkout uses whatever strategy has been set.
func (sc *ShoppingCart) Checkout() {
	if sc.paymentStrategy == nil {
		fmt.Println("Please select a payment strategy.")
		return
	}
	// This is the dynamic dispatch!
	// It calls the Pay() method on whichever concrete struct was provided.
	result := sc.paymentStrategy.Pay(sc.totalAmount)
	fmt.Println(result)
}

func main() {
	// Create the concrete strategy objects.
	creditCard := &CreditCardStrategy{CardNumber: "1234-5678-9876-5432", Owner: "John Doe"}
	payPal := &PayPalStrategy{Email: "john.doe@example.com"}

	// Create the context object.
	cart := &ShoppingCart{totalAmount: 150.75}

	// Use the credit card strategy first.
	fmt.Println("--- Checking out with Credit Card ---")
	cart.SetPaymentStrategy(creditCard)
	cart.Checkout()

	fmt.Println("\n--- Customer changes their mind, now checking out with PayPal ---")
	// Now, swap the strategy on the SAME cart object.
	cart.SetPaymentStrategy(payPal)
	cart.Checkout()
}
