package main

import (
	"fmt"
)

// Coffee is the base component interface
type Coffee interface {
	GetCost() float64
	GetDescription() string
}

// BasicCoffee is the concrete component
type BasicCoffee struct{}

func (b *BasicCoffee) GetCost() float64 {
	return 2.0
}

func (b *BasicCoffee) GetDescription() string {
	return "Basic Coffee"
}

// MilkDecorator is a decorator
type MilkDecorator struct {
	Coffee Coffee
}

func (m *MilkDecorator) GetCost() float64 {
	return m.Coffee.GetCost() + 0.5
}

func (m *MilkDecorator) GetDescription() string {
	return m.Coffee.GetDescription() + ", Milk"
}

// SugarDecorator is another decorator
type SugarDecorator struct {
	Coffee Coffee
}

func (s *SugarDecorator) GetCost() float64 {
	return s.Coffee.GetCost() + 0.3
}

func (s *SugarDecorator) GetDescription() string {
	return s.Coffee.GetDescription() + ", Sugar"
}

// MAIN
func main() {
	var myCoffee Coffee = &BasicCoffee{}
	fmt.Println(myCoffee.GetDescription(), "$", myCoffee.GetCost())

	// Add Milk
	myCoffee = &MilkDecorator{Coffee: myCoffee}
	fmt.Println(myCoffee.GetDescription(), "$", myCoffee.GetCost())

	// Add Sugar
	myCoffee = &SugarDecorator{Coffee: myCoffee}
	fmt.Println(myCoffee.GetDescription(), "$", myCoffee.GetCost())
}
