// The ability of different struct or class to be treated as instances of the same class through inheritance.

package main

import "fmt"

type Shape interface {
	Draw()
}

type Circle struct{}

func (c *Circle) Draw() {
	fmt.Println("Drawing a Circle...")
}

type Square struct{}

func (s *Square) Draw() {
	fmt.Println("Drawing a Square...")
}

func main() {
	var shape Shape

	shape = &Circle{}
	shape.Draw()

	shape = &Square{}
	shape.Draw()

}

// Key Takeaways = Polymorphism allows different objects (Circle, Sqaure) to respond to the same  method (Draw), but behave differently.
