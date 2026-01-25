package main

import "fmt"

// Prototype Interface
type Prototype interface {
	Clone() Prototype
}

// Concrete type
type Shape struct {
	Type  string
	Color string
}

func (s *Shape) Clone() Prototype {
	return &Shape{Type: s.Type, Color: s.Color}
}

func main() {
	shape1 := &Shape{Type: "Circle", Color: "Red"}
	shape2 := shape1.Clone().(*Shape)
	fmt.Println(shape1, shape2)
}
