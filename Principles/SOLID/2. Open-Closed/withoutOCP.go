// A module should be open for extension code but closed for modification!
package main

import "fmt"

type PaymentService struct{}

func (p *PaymentService) Pay(method string, amount float64) {
	if method == "Credit-card" {
		fmt.Println("Paying with Credit Card:", amount)
	} else if method == "Paypal" {
		fmt.Println("Paying with Paypal:", amount)
	}
}

func main() {
	pay := &PaymentService{}
	pay.Pay("Credit-card", 200.00)
	// pay.Pay("Paypal", 200.00)
}
