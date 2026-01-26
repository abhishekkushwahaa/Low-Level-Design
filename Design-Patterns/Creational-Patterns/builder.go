package main

type Computer struct {
	CPU string
	RAM string
}

// Builder
type ComputerBuilder struct {
	computer Computer
}

// 1. Simpler Builder (Fluent API)
// Benefits: a)- Readable b)- Flexible c)- Avoids large constructor
// Problems: a)- No strict construction order b)- Client controls all steps
func NewComputerBuilder() *ComputerBuilder {
	return &ComputerBuilder{}
}

func (b *ComputerBuilder) SetCPU(cpu string) *ComputerBuilder {
	b.computer.CPU = cpu
	return b
}

func (b *ComputerBuilder) SetRAM(ram string) *ComputerBuilder {
	b.computer.RAM = ram
	return b
}

func (b *ComputerBuilder) Build() Computer {
	return b.computer
}
