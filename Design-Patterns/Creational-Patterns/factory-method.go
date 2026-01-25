package main

import "fmt"

// Product interface
type Notification interface {
	Send(msg string)
}

// Concrete products
type EmailNotifications struct{}
type WhatsAppNotifications struct{}

func (e EmailNotifications) Send(msg string) {
	fmt.Println("Sending notifications on Email", msg)
}

func (w WhatsAppNotifications) Send(msg string) {
	fmt.Println("Sending notifications on WhatsApp", msg)
}

// 1. Basic Factory Method (switch-based)
// Benefits: Centralized object creation
// Problems: Must modify factory for every new type
func GetNotification(notificationType string) Notification {
	switch notificationType {
	case "email":
		return EmailNotifications{}
	case "whatsapp":
		return WhatsAppNotifications{}
	default:
		return nil
	}
}
