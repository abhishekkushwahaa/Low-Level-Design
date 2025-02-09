// The concept of bundling the data and the methods.

package main

import "fmt"

type Car struct {
	brand  string
	engine string
}

func NewCar(brand string, engine string) *Car {
	return &Car{brand: brand, engine: engine}
}

func (c *Car) StartCar() {
	fmt.Println("Start:", c.brand, "Car:", c.engine)
}

func main() {
	car := NewCar("Tesla", "Electric")
	car.StartCar()
}

// Key Takeaways = Encapsulation hides the internal details (like the engine), and only exposes relevant action (like Start Car)
