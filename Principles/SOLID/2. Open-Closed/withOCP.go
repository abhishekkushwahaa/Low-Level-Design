// A module should be open for extension code but closed for modification!
package main

import "fmt"

type PaymentProcess interface {
	Process(amount float64)
}

type CreditCardProcess struct{}

func (c *CreditCardProcess) Process(amount float64) {
	fmt.Println("Paying with Credit-card:", amount)
}

type PaypalProcess struct{}

func (p *PaypalProcess) Process(amount float64) {
	fmt.Println("Paying with Paypal:", amount)
}

type PaymentService struct{}

func (p *PaymentService) Pay(process PaymentProcess, amount float64) {
	process.Process(amount)
}

func main() {
	pay := &PaymentService{}
	pay.Pay(&CreditCardProcess{}, 200.00)
}
