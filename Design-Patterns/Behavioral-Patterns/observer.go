package main

import (
	"fmt"
)

// Subscriber Interface
type Subscriber interface {
	Update(eventType string, data string)
}

// Publisher Struct
type Publisher struct {
	Subscribers map[string][]Subscriber
}

func NewPublisher() *Publisher {
	return &Publisher{Subscribers: make(map[string][]Subscriber)}
}

// Subscribe
func (p *Publisher) Subscriber(eventType string, subscriber Subscriber) {
	p.Subscribers[eventType] = append(p.Subscribers[eventType], subscriber)
}

// Unsubscribe
func (p *Publisher) Unsubscribe(eventType string, subscriber Subscriber) {
	subscribers := p.Subscribers[eventType]
	for i, sub := range subscribers {
		if sub == subscriber {
			p.Subscribers[eventType] = append(subscribers[:i], subscribers[i+1:]...)
		}
	}
}

// Notify
func (p *Publisher) Notify(eventType string, data string) {
	for _, subscriber := range p.Subscribers[eventType] {
		subscriber.Update(eventType, data)
	}
}

// Concrete Subscriber 1 - Logging
type LoggingSubscriber struct {
	LogFile string
}

func (l *LoggingSubscriber) Update(eventType string, data string) {
	fmt.Println("Logging:", l.LogFile, eventType, data)
}

// Concrete Subscriber 2 - Email Alert
type EmailSubscriber struct {
	email string
}

func (e *EmailSubscriber) Update(eventType string, data string) {
	fmt.Println("Send Email:", e.email, eventType, data)
}

// Main
func main() {
	publisher := NewPublisher()

	// Subscribers
	loggingsub := &LoggingSubscriber{LogFile: "path/log.txt"}
	emailsub := &EmailSubscriber{email: "abc@gmail.com"}

	// Subscribe
	publisher.Subscriber("file-open", loggingsub)
	publisher.Subscriber("file-save", emailsub)

	// Notify Events
	publisher.Notify("file-open", "log.txt")
	publisher.Notify("file-save", "log.txt")
}
