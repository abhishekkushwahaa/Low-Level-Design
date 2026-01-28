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

// 2. Builder with Director
// Benefits: a)- Separates construction from representation b)- Same process can build different objects
// Problems: a)- More complex b)- Extra abstraction
type ComputerBuilderDirector interface {
	BuildCPU()
	BuildRAM()
	GetComputer() Computer
}

// Concrete Builder
type GamingComputerBuilder struct {
	computer Computer
}

func (g *GamingComputerBuilder) BuildCPU() {
	g.computer.CPU = "Intel i9"
}

func (g *GamingComputerBuilder) BuildRAM() {
	g.computer.RAM = "RTX 4090"
}

func (g *GamingComputerBuilder) GetComputer() Computer {
	return g.computer
}

// Director
type Director struct {
	builder ComputerBuilderDirector
}

func NewDirector(b ComputerBuilderDirector) *Director {
	return &Director{builder: b}
}

func (d *Director) Construct() Computer {
	d.builder.BuildCPU()
	d.builder.BuildRAM()
	return d.builder.GetComputer()
}
