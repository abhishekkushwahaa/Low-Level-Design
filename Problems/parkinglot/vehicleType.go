package main

// VehicleType Enum
type VehicleType int

const (
	CAR VehicleType = iota
	TRUCK
	MOTORCYCLE
)

// Convert enum to string
func (v VehicleType) String() string {
	return [...]string{"CAR", "TRUCK", "MOTORCYCLE"}[v]
}
