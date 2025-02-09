/**
Method Overloading (Not in Go) = It allows a class to have multiple methods with same name but different parameters. However Does not support this directly, but the concept is still widely used in many other language.

Method Overriding = IT is the ability of a subclass or derived class to provide a specific implementation of a method that is already defined in its super class or base class
**/

package main

import "fmt"

// Method Overloading - Go Alternative (Using Variadic Functions)
// Print method can handle multiple data types by using empty interface
func Print(values ...interface{}) {
	for _, v := range values {
		fmt.Println(v)
	}
}

// Key Takeaways = Go does not support Method Overloading but allows flexibility using variadic function and interface

// Method Overriding = Think of a Car and a Sports Car
type Vehicle struct{}

func (v *Vehicle) Drive() {
	fmt.Println("Deriving the vehicle...")
}

type SportCar struct {
	Vehicle
}

func (s *SportCar) Drive() {
	fmt.Println("Deriving the Sports Car at high speed...")
}

func main() {
	Print("Hello", 42, 1.2) // Accepts different types (string, int, float)

	vehicle := &Vehicle{}
	sportsCar := &SportCar{}
	vehicle.Drive()   // Calls Vehicle Drive
	sportsCar.Drive() // Calls SportsCar's Drive (Overriding Vehicle Method)
}
