package main

import "fmt"

// Component Interface
type Coffee interface {
	Cost() float64
	Description() string
}

// Concrete component (Basic Coffee)
type BasicCoffee struct{}

func (b *BasicCoffee) Cost() float64 {
	return 5.0 // Base Cost of Black Coffee
}

func (b *BasicCoffee) Description() string {
	return "Black Coffee"
}

// Decorator base struct
type CoffeeDecorator struct {
	coffee Coffee
}

// Concrete Decorator
// Milk Add-On
type Milk struct {
	CoffeeDecorator
}

func (m *Milk) Cost() float64 {
	return m.coffee.Cost() + 1.5 // Milk Cost = 1.5
}

func (m *Milk) Description() string {
	return m.coffee.Description() + " With Milk"
}

// Sugar Add-On
type Sugar struct {
	CoffeeDecorator
}

func (s *Sugar) Cost() float64 {
	return s.coffee.Cost() + 0.5 // Sugar Cost = 0.5
}

func (s *Sugar) Description() string {
	return s.coffee.Description() + ", Sugar"
}

// Whipped Cream Add-On
type WhippedCream struct {
	CoffeeDecorator
}

func (w *WhippedCream) Cost() float64 {
	return w.coffee.Cost() + 2.5 // Whipped Cream Cost = 2.5
}

func (w *WhippedCream) Description() string {
	return w.coffee.Description() + " & Whipped Cream!"
}

// Client Code
func main() {
	// Start with basic black coffee
	coffee := &BasicCoffee{}
	fmt.Println(coffee.Description(), "-", coffee.Cost())

	// Add Milk
	coffeeWithMilk := &Milk{CoffeeDecorator{coffee}}
	fmt.Println(coffeeWithMilk.Description(), "-", coffeeWithMilk.Cost())

	// Sugar Added
	coffeeWithMilkandSugar := &Sugar{CoffeeDecorator{coffeeWithMilk}}
	fmt.Println(coffeeWithMilkandSugar.Description(), "-", coffeeWithMilkandSugar.Cost())

	// Add Whipped Cream
	coffeeWithAll := &WhippedCream{CoffeeDecorator{coffeeWithMilkandSugar}}
	fmt.Println(coffeeWithAll.Description(), "-", coffeeWithAll.Cost())
}
