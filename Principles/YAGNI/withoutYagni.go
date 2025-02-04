// YAGNI = You Aren't Gonna Need It

package main

import "fmt"

type Product struct {
	Name  string
	Price float64
}

func getProductName() {
	name := Product{Name: "MacBook!"}
	fmt.Println("Product Name -", name.Name)
}

// Unnecessary discount logic (not needed yet)
func (p *Product) getDiscountedPrice() float64 {
	return p.Price * 0.9
}

func main() {
	getProductName()
}
