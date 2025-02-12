package main

type Spot struct {
	ID          int
	VehicleType VehicleType
	IsOccupied  bool
}

// Constructor for Spots
func NewSpot(id int, vType VehicleType) *Spot {
	return &Spot{ID: id, VehicleType: vType, IsOccupied: false}
}

// Assign a vehicle to this spot
func (s *Spot) AssignVehicle() {
	s.IsOccupied = true
}

// Released a spot when Vehicle the leaves
func (s *Spot) ReleasedVehicle() {
	s.IsOccupied = false
}
