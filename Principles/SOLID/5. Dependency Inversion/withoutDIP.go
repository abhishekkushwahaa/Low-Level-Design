// High-level modules should not depend on low-level modules.
package main

import "fmt"

// Low-level module
type NotificationService struct{}

func (n *NotificationService) Notify(email string) {
	fmt.Println("Send email:", email)
}

// High-level module directly depended on Low-level module
func main() {
	notification := &NotificationService{}
	notification.Notify("abc@gmail.com")
}
