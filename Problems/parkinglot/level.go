package main

type Level struct {
	ID    int
	Spots []*Spot
}

// Constructor for Level with different types of spots
func NewLevel(id int, spotCounts map[VehicleType]int) *Level {
	level := &Level{ID: id}
	spotId := 1
	for vType, count := range spotCounts {
		for i := 0; i < count; i++ {
			level.Spots = append(level.Spots, NewSpot(spotId, vType))
			spotId++
		}
	}
	return level
}

// Find an available parking spot for a given vehicle type
func (l *Level) FindSpot(vType VehicleType) *Spot {
	for _, spot := range l.Spots {
		if spot.VehicleType == vType && !spot.IsOccupied {
			return spot
		}
	}
	return nil
}

// Park a vehicle in this level
func (l *Level) ParkVehicle(vehicle *Vehicle) *Spot {
	spot := l.FindSpot(vehicle.Type)
	if spot != nil {
		spot.AssignVehicle()
		return spot
	}
	return nil
}

// Free a spot when a vehicle leaves
func (l *Level) LeaveVehicle(spotID int) {
	for _, spot := range l.Spots {
		if spot.ID == spotID {
			spot.ReleasedVehicle()
			break
		}
	}
}
