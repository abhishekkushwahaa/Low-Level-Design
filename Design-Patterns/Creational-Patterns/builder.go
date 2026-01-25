package main

import "fmt"

// Product: Car
type Car struct {
	Seat         int
	Engine       string
	TripComputer bool
	GPS          bool
}

// Builder Interface: Define construction steps
type Builder interface {
	SetSeats(seats int)
	SetEngine(engine string)
	SetTripComputer(tripcomputer bool)
	SetGPS(gps bool)
	GetCar() *Car
}

// Concrete Builder: Car Builder

type CarBuilder struct {
	Car *Car
}

func NewCarBuilder() *CarBuilder {
	return &CarBuilder{Car: &Car{}}
}

func (b *CarBuilder) SetSeats(seats int) {
	b.Car.Seat = seats
}

func (b *CarBuilder) SetEngine(engine string) {
	b.Car.Engine = engine
}
func (b *CarBuilder) SetTripComputer(tripcomputer bool) {
	b.Car.TripComputer = tripcomputer
}
func (b *CarBuilder) SetGPS(gps bool) {
	b.Car.GPS = gps
}

func (b *CarBuilder) GetCar() *Car {
	return b.Car
}

// Director: Guides the Construction process
type Director struct{}

func (d *Director) ConstructSportsCar(builder Builder) {
	builder.SetSeats(2)
	builder.SetEngine("V8")
	builder.SetTripComputer(true)
	builder.SetGPS(true)
}

func (d *Director) ConstructSUV(builder Builder) {
	builder.SetSeats(5)
	builder.SetEngine("V5")
	builder.SetTripComputer(true)
	builder.SetGPS(true)
}

func main() {
	director := &Director{}

	// Client creates a builder and passes it to the Director
	builder := NewCarBuilder()

	// Director constructs a Sports Car
	director.ConstructSportsCar(builder)
	car := builder.GetCar()
	fmt.Printf("Sports Car: %+v\n", car)

	// Director constructs an SUV Car
	builder = NewCarBuilder()
	director.ConstructSUV(builder)
	suv := builder.GetCar()
	fmt.Printf("SUV Car: %+v\n", suv)
}
