// High-level modules should not depend on low-level modules. Both should depend on abstractions.
package main

import "fmt"

type Notifier interface {
	Notify(Message string)
}

type EmailNotifier struct{}

func (e *EmailNotifier) Notify(message string) {
	fmt.Println("Sending email:", message)
}

type SMSNotifier struct{}

func (e *SMSNotifier) Notify(message string) {
	fmt.Println("Sending sms:", message)
}

// High-level module depends on the abstraction (Notifier interface)
type NotificationService struct {
	notifier Notifier
}

func (n *NotificationService) SendNotification(message string) {
	n.notifier.Notify(message)
}

func main() {

	// Using EmailNotifier
	emailNotifier := &EmailNotifier{}
	emailNotifierService := &NotificationService{notifier: emailNotifier}
	emailNotifierService.SendNotification("Sent Email!")

	// Using SMSNotifier
	smsNotifier := &SMSNotifier{}
	smsNotificationService := &NotificationService{notifier: smsNotifier}
	smsNotificationService.SendNotification("Sent SMS!")
}
