// Subtypes must be substitutable for their base types without altering program correctness.

package main

import "fmt"

type Vehicle interface {
	Drive()
}

type Car struct{}

func (c *Car) Drive() {
	fmt.Println("Driving a Car")
}

type Bus struct{}

func (b *Bus) Drive() {
	fmt.Println("Driving a Bus")
}

type Plane struct{}

func (p Plane) Drive() {
	fmt.Println("Flying a Plane")
}

func TestDrive(v Vehicle) {
	v.Drive()
}

func main() {
	car := Car{}
	bus := Bus{}
	plane := Plane{}
	TestDrive(&car)
	TestDrive(&bus)
	TestDrive(&plane)
}

// Key Takeaways = Any Subclass should behave as expected when used in place of its parent class.
