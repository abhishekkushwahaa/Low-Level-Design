// When we build complex object by combining simpler objects
package main

import "fmt"

type Engine struct{}

func (e *Engine) Start() {
	fmt.Println("Engine is Starting...")
}

type Car struct {
	engine Engine
}

func (c *Car) StartCar() {
	c.engine.Start()
	fmt.Println("Car is Starting...")
}

func main() {
	car := Car{engine: Engine{}}
	car.StartCar()
}

// Key Takeaways = Composition allows one object to include another.
