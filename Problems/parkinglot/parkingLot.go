package main

import "fmt"

type ParkingLot struct {
	Levels []*Level
}

// Constructor for ParkingLot
func NewParkingLot(levelCount int, spotConfig map[VehicleType]int) *ParkingLot {
	parkingLot := &ParkingLot{}

	for i := 0; i < levelCount; i++ {
		parkingLot.Levels = append(parkingLot.Levels, NewLevel(i, spotConfig))
	}

	return parkingLot
}

// Park a vehicle in the first available spot
func (p *ParkingLot) ParkVehicle(vehicle *Vehicle) *Spot {
	for _, level := range p.Levels {
		spot := level.ParkVehicle(vehicle)
		if spot != nil {
			fmt.Printf("Vehicle parked at Level %d, Spot %d\n", level.ID, spot.ID)
			return spot
		}
	}
	fmt.Println("No available spot")
	return nil
}

// Free a parking spot when a vehicle leaves
func (p *ParkingLot) LeaveVehicle(levelId int, spotId int) {
	for _, level := range p.Levels {
		if levelId == level.ID {
			level.LeaveVehicle(spotId)
			fmt.Printf("Vehicle left from Level %d, Spot %d\n", level.ID, spotId)
			return
		}
	}
	fmt.Println("Spot not found")
}

// Display available spots
func (p *ParkingLot) GetAvailability() {
	for _, level := range p.Levels {
		available := 0
		for _, spot := range level.Spots {
			if !spot.IsOccupied {
				available++
			}
		}
		fmt.Printf("Level %d: %d available spots\n", level.ID, available)
	}
}
