package main

type Vehicle struct {
	Type VehicleType
}

// Constructor New Vehicle
func NewVehicle(vType VehicleType) *Vehicle {
	return &Vehicle{Type: vType}
}
