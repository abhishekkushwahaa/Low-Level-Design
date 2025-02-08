// Subtypes must be substitutable for their base types without altering program correctness.

package main

import "fmt"

// Base Type
type Car struct{}

func (c Car) Drive() {
	fmt.Println("Driving a Car")
}

// Subtype 1
type Bus struct{}

func (b *Bus) Drive() {
	fmt.Println("Driving a Bus")
}

// Subtype 2 (Violates LSP)
type Plane struct{}

func (p *Plane) Drive() {
	panic("Planes don't Drive") // This breaks the expected behavior
}

func TestDrive(vehicle interface{}) {
	car := vehicle.(Car)
	car.Drive()
}

func main() {
	car := Car{}
	bus := Bus{}
	plane := Plane{}

	// This works fine
	TestDrive(car)

	// This will cause a runtime panic because Bus cannot be asserted to Car
	TestDrive(bus)
	TestDrive(plane)
}
