package main

import "fmt"

// Product Interface for transport
type Transport interface {
	Deliver()
}

// Truck Struct (Concrete Product)
type Truck struct{}

func (t *Truck) Deliver() {
	fmt.Println("Delivering box via land in the Truck!")
}

// Ship Struct (Concrete Product)
type Ship struct{}

func (s *Ship) Deliver() {
	fmt.Println("Delivering box via sea in the Ship!")
}

// Logistics Interface (Creator)
type Logistics interface {
	CreateTransport() Transport
}

// Road Logistics struct (Concrete Creator)
type RoadLogistics struct{}

func (r *RoadLogistics) CreateTransport() Transport {
	return &Truck{}
}

// Sea Logistics struct (Concrete Creator)
type SeaLogistics struct{}

func (s *SeaLogistics) CreateTransport() Transport {
	return &Ship{}
}

// Client Code
func planDelivery(logistics Logistics) {
	transport := logistics.CreateTransport() // Factory Method
	transport.Deliver()                      // Polymorphic Call
}

// func main() {
// 	// Road Delivery
// 	roadLogistics := RoadLogistics{}
// 	planDelivery(&roadLogistics)

// 	// Sea Delivery
// 	seaLogistics := SeaLogistics{}
// 	planDelivery(&seaLogistics)
// }
