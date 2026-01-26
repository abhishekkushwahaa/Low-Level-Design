package main

type Computer struct {
	CPU string
	RAM string
	SSD string
	GPU string
}

// Builder
type ComputerBuilder struct {
	computer Computer
}
