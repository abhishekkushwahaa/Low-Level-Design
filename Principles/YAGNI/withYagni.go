// YAGNI = You Aren't Gonna Need It

package main

import "fmt"

type Product struct {
	Name string
}

func getProductName() {
	name := Product{Name: "MacBook!"}
	fmt.Println("Product Name -", name.Name)
}

func main() {
	getProductName()
}
