package main

import "fmt"

// Structure :-
// Abstract Products → Interfaces (Payment, Refund)
// Concrete Products → UPIFamily, CardFamily
// Abstract Factory → PaymentFactory
// Concrete Factories → UPIFactory, CardFactory
// Client → Uses only interfaces

// Abstract Products
type Payment interface {
	Pay(amount int)
}

type Refund interface {
	Process(amount int)
}

// Concrete Products - UPI Family
type UPIPayment struct{}
type UPIRefund struct{}

func (u UPIPayment) Pay(amount int) {
	fmt.Println("UPI Payment:", amount)
}

func (u UPIRefund) Process(amount int) {
	fmt.Println("UPI Refund:", amount)
}

// Concrete Products - Card Family
type CardPayment struct{}
type CardRefund struct{}

func (c CardPayment) Pay(amount int) {
	fmt.Println("Card Patyment:", amount)
}

func (c CardRefund) Process(amount int) {
	fmt.Println("Card Refund:", amount)
}

// Abstract Factory
type PaymentFactory interface {
	CreatePayment() Payment
	CreateRefund() Refund
}

// Concrete Factories
type UPIFactory struct{}
type CardFactory struct{}

func (u UPIFactory) CreatePayment() Payment {
	return UPIPayment{}
}

func (u UPIFactory) CreateRefund() Refund {
	return UPIRefund{}
}

func (c CardFactory) CreatePayment() Payment {
	return CardPayment{}
}

func (c CardFactory) CreateRefund() Refund {
	return CardRefund{}
}

// Client
// func main() {
// 	var factory PaymentFactory

// 	factory = UPIFactory{}
// 	payment := factory.CreatePayment()
// 	refund := factory.CreateRefund()

// 	payment.Pay(1000)
// 	refund.Process(500)

// 	factory = CardFactory{}
// 	payment1 := factory.CreatePayment()
// 	refund1 := factory.CreateRefund()

// 	payment1.Pay(2000)
// 	refund1.Process(1000)
// }
