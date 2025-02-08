// Should not be forced to implemented interfaces it doesn't use.
package main

import "fmt"

type Coder interface {
	Code()
}

type Designer interface {
	Design()
}

type Tester interface {
	Test()
}

type Developer struct{}

func (d *Developer) Code() {
	fmt.Println("Coding...")
}

type GraphicDesigner struct{}

func (g *GraphicDesigner) Designer() {
	fmt.Println("Desinging...")
}

type TesterEngineer struct{}

func (t *TesterEngineer) Tester() {
	fmt.Println("Testing...")
}

func main() {
	dev := &Developer{}
	dev.Code()

	design := &GraphicDesigner{}
	design.Designer()

	test := &TesterEngineer{}
	test.Tester()
}
