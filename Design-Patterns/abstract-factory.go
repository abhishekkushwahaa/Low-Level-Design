package main

import "fmt"

// Abstract Product A
type Chair interface {
	SitOn()
}

// Abstract Product B
type Table interface {
	Use()
}

// Concrete Product A1
type VictorianChair struct{}

func (c VictorianChair) SitOn() {
	fmt.Println("Sitting on a V-Chair!")
}

// Concrete Product A2
type ModernChair struct{}

func (c ModernChair) SitOn() {
	fmt.Println("Sitting on a M-Chair!")
}

// Concrete Product B1
type VictorianTable struct{}

func (t VictorianTable) Use() {
	fmt.Println("Using a V-Table!")
}

// Concrete Product B2
type ModernTable struct{}

func (t ModernTable) Use() {
	fmt.Println("Using a M-Table!")
}

// Abstract Factory
type FurnitureFactory interface {
	CreateChair() Chair
	CreateTable() Table
}

// Concrete Factory 1
type VFurnitureFactory struct{}

func (f VFurnitureFactory) CreateChair() Chair {
	return VictorianChair{}
}

func (f VFurnitureFactory) CreateTable() Table {
	return VictorianTable{}
}

// Concrete Factory 2
type MFurnitureFactory struct{}

func (f MFurnitureFactory) CreateChair() Chair {
	return ModernChair{}
}

func (f MFurnitureFactory) CreateTable() Table {
	return ModernTable{}
}

// Client Code
func ClientCode(factory FurnitureFactory) {
	chair := factory.CreateChair()
	table := factory.CreateTable()
	chair.SitOn()
	table.Use()
}

func main() {
	var factory FurnitureFactory

	// Use Victorian Furniture Factory
	factory = VFurnitureFactory{}
	ClientCode(factory)

	//Use Modern Furniture Factory
	factory = MFurnitureFactory{}
	ClientCode(factory)
}
