package main

// Component interface
type Coffee interface {
	Cost() int
	Description() string
}

// Concrete Component
type SimpleCoffee struct{}

func (s *SimpleCoffee) Cost() int {
	return 100
}

func (s *SimpleCoffee) Description() string {
	return "Simple Coffee"
}

/*
1. Basic Decorator
Benefits:
a)- Adds new behavior without modifying existing code
b)- Follows Open/Closed Principle
c)- Flexible alternative to subclassing
Problems:
a)- Many small structs created
b)- Can be hard to debug
*/

// Base Decorator
type CoffeeDecorator struct {
	coffee Coffee
}

func (d *CoffeeDecorator) Cost() int {
	return d.coffee.Cost()
}

func (d *CoffeeDecorator) Description() string {
	return d.coffee.Description()
}

// Concrete Decorator - Milk
type MilkDecorator struct {
	CoffeeDecorator
}

func NewMilkDecorator(c Coffee) Coffee {
	return &MilkDecorator{CoffeeDecorator{coffee: c}}
}

func (m *MilkDecorator) Cost() int {
	return m.coffee.Cost() + 20
}

func (m *MilkDecorator) Description() string {
	return m.coffee.Description() + ", Milk"
}

/*
2. Multiple Decorators (Chaining)
Benefits:
a)- Combine multiple features dynamically
b)- No need for multiple subclasses
c)- Behavior added at runtime
Problems:
a)- Order of decorators matters
b)- Harder to trace execution flow
*/

// Concrete Decorator - Sugar
type SugarDecorator struct {
	CoffeeDecorator
}

func NewSugarDecorator(c Coffee) Coffee {
	return &SugarDecorator{CoffeeDecorator{coffee: c}}
}

func (s *SugarDecorator) Cost() int {
	return s.coffee.Cost() + 10
}

func (s *SugarDecorator) Description() string {
	return s.coffee.Description() + ", Sugar"
}

/*
3. Decorator Registry (Dynamic Decoration)
Benefits:
a)- Centralized decorator management
b)- Easy to add/remove features dynamically
c)- Scalable for many decorators
Problems:
a)- Registry complexity
b)- Risk of wrong decorator key usage
*/

var decoratorRegistry = map[string]func(Coffee) Coffee{}

func RegisterDecorator(key string, decorator func(Coffee) Coffee) {
	decoratorRegistry[key] = decorator
}

func GetDecoratedCoffee(key string, coffee Coffee) Coffee {
	if decorator, ok := decoratorRegistry[key]; ok {
		return decorator(coffee)
	}
	return coffee
}
