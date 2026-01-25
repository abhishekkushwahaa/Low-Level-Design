package main

import "fmt"

// Strategy Interface
type RouteStrategy interface {
	BuildRoute(start, destination string)
}

// Concrete Strategy 1 - Driving
type Driving struct{}

func (d *Driving) BuildRoute(start, destination string) {
	fmt.Printf("Driving route \n", start, destination)
}

// Concrete Strategy 2 - Walking
type Walking struct{}

func (s *Walking) BuildRoute(start, destination string) {
	fmt.Print("Walking route \n", start, destination)
}

// Context - Navigator
type Navigator struct {
	Strategy RouteStrategy
}

func (n *Navigator) SetStrategy(strategy RouteStrategy) {
	n.Strategy = strategy
}

func (n *Navigator) BuildRoute(start, destination string) {
	n.Strategy.BuildRoute(start, destination)
}

func main() {
	navigator := &Navigator{}
	// Set Driving Strategy
	navigator.SetStrategy(&Driving{})
	navigator.BuildRoute("Home -", " Office")

	// Change to Walking Strategy
	navigator.SetStrategy(&Walking{})
	navigator.BuildRoute("Home -", " Park")
}
