package main

func main() {
	// Define parking spot configuration per level
	spotConfig := map[VehicleType]int{
		CAR:        3,
		TRUCK:      2,
		MOTORCYCLE: 2,
	}

	// Create a Parking Lot with 2 levels
	parkingLot := NewParkingLot(2, spotConfig)

	// Create Vehicle
	car1 := NewVehicle(CAR)
	truck1 := NewVehicle(TRUCK)
	motorcycle1 := NewVehicle(MOTORCYCLE)

	// Park Vehicles
	spot1 := parkingLot.ParkVehicle(car1)
	_ = parkingLot.ParkVehicle(truck1)
	_ = parkingLot.ParkVehicle(motorcycle1)

	// Check Availability
	parkingLot.GetAvailability()

	// Released a Spot
	if spot1 != nil {
		parkingLot.LeaveVehicle(1, spot1.ID)
	}

	// Check Availability Again
	parkingLot.GetAvailability()
}
