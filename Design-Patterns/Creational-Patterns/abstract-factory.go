package main

// Structure :-
// Abstract Products → Interfaces (Payment, Refund)
// Concrete Products → UPIFamily, UPIFamily
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

// Concrete Products - UPI Family

// Abstract Factory

// Concrete Factories

// Client
