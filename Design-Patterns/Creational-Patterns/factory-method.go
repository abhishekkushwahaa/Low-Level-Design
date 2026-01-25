package main

import "fmt"

// Product interface
type Notification interface {
	Send(msg string)
}

// Concrete products
type EmailNotification struct{}
type WhatsAppNotification struct{}

func (e EmailNotification) Send(msg string) {
	fmt.Println("Sending notification on Email:", msg)
}

func (w WhatsAppNotification) Send(msg string) {
	fmt.Println("Sending notification on WhatsApp:", msg)
}

// 1. Basic Factory Method (switch-based)
// Benefits: Centralized object creation
// Problems: Must modify factory for every new type
func GetNotification(notificationType string) Notification {
	switch notificationType {
	case "email":
		return EmailNotification{}
	case "whatsapp":
		return WhatsAppNotification{}
	default:
		return nil
	}
}

// 2. Factory using Map (Extensible)
// Benefits: Easy to add new types without modifying logic
var factoryMap = map[string]func() Notification{
	"email":    func() Notification { return EmailNotification{} },
	"whatsapp": func() Notification { return WhatsAppNotification{} },
}

func GetNotificationMap(notificationType string) Notification {
	if factory, ok := factoryMap[notificationType]; ok {
		return factory()
	}
	return nil
}
