// Should not be forced to implemented interfaces it doesn't use.
package main

import "fmt"

type Worker interface {
	Code()
	Design()
	Test()
}

// Developer implements the Worker interface
type Developer struct{}

func (d *Developer) Code() {
	fmt.Println("Coding...")
}

// Developer is forced to implement Design, even though it might not be their responsibility
func (d *Developer) Design() {
	fmt.Println("Designing...")
}

// Developer is forced to implement Test, even though it might not be their responsibility
func (d *Developer) Test() {
	fmt.Println("Testing...")
}

func main() {
	dev := &Developer{}
	dev.Code()
	dev.Design()
	dev.Test()
}
