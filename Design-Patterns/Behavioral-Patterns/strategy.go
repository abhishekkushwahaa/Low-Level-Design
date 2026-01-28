package main

import "fmt"

// Strategy interface
type PaymentStrategy interface {
	Pay(amount int)
}

// Concrete Strategies
type UPIStrategy struct{}
type CardStrategy struct{}
type CashStrategy struct{}

func (u *UPIStrategy) Pay(amount int) {
	fmt.Println("Paid via UPI:", amount)
}

func (c *CardStrategy) Pay(amount int) {
	fmt.Println("Paid via Card:", amount)
}

func (c *CashStrategy) Pay(amount int) {
	fmt.Println("Paid via Cash:", amount)
}

/*
1. Basic Strategy (Interface-based)
Benefits:
a)- Eliminates if-else conditions
b)- Follows Open/Closed Principle
c)- Easy to add new strategies
Problems:
a)- More objects created
b)- Client must choose strategy
*/

func ProcessPayment(strategy PaymentStrategy, amount int) {
	strategy.Pay(amount)
}

/*
2. Strategy with Context
Benefits:
a)- Encapsulates strategy usage
b)- Client code simplified
c)- Strategy can be changed at runtime
Problems:
a)- Extra abstraction
b)- Slight complexity increase
*/

// Context
type PaymentContext struct {
	strategy PaymentStrategy
}

func NewPaymentContext(strategy PaymentStrategy) *PaymentContext {
	return &PaymentContext{strategy: strategy}
}

func (p *PaymentContext) SetStrategy(strategy PaymentStrategy) {
	p.strategy = strategy
}

func (p *PaymentContext) Pay(amount int) {
	p.strategy.Pay(amount)
}

/*
3. Strategy Registry (Dynamic Selection)
Benefits:
a)- Centralized strategy management
b)- Easy runtime selection
c)- Scalable for many strategies
Problems:
a)- Registry maintenance
b)- Risk of wrong key usage
*/
var strategyRegistry = map[string]PaymentStrategy{}

func RegisterStrategy(key string, strategy PaymentStrategy) {
	strategyRegistry[key] = strategy
}

func GetStrategy(key string) PaymentStrategy {
	if strategy, ok := strategyRegistry[key]; ok {
		return strategy
	}
	return nil
}
