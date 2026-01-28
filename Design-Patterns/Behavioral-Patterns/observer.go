package main

import "fmt"

// Observer interface
type Observer interface {
	Update(msg string)
}

// Subject interface
type Subject interface {
	Register(o Observer)
	Remove(o Observer)
	Notify(msg string)
}

/*
1. Basic Observer (Interface-based)
Benefits:
a)- Loose coupling between subject and observers
b)- Easy to add new observers
c)- Follows Open/Closed Principle
Problems:
a)- All observers get notified even if not needed
b)- Notification order is not guaranteed
*/

// Concrete Observer
type UserObserver struct {
	name string
}

func (u *UserObserver) Update(msg string) {
	fmt.Println("User", u.name, "received:", msg)
}

/*
2. Observer with Subject (Concrete Subject + Observers)
Benefits:
a)- Centralized event management
b)- Automatic notification on state change
c)- Decouples publisher from subscribers
Problems:
a)- Subject must manage observer list
b)- Potential memory leaks if observers not removed
*/

// Concrete Subject
type NotificationService struct {
	observers []Observer
}

func (n *NotificationService) Register(o Observer) {
	n.observers = append(n.observers, o)
}

func (n *NotificationService) Remove(o Observer) {
	for i, observer := range n.observers {
		if observer == o {
			n.observers = append(n.observers[:i], n.observers[i+1:]...)
			break
		}
	}
}

func (n *NotificationService) Notify(msg string) {
	for _, observer := range n.observers {
		observer.Update(msg)
	}
}

/*
3. Observer Registry (Multiple Subscribers)
Benefits:
a)- Centralized observer storage
b)- Dynamic subscribe/unsubscribe
c)- Useful for event-driven systems
Problems:
a)- Registry management complexity
b)- Risk of stale observers
*/
var observerRegistry = map[string][]Observer{}

func RegisterObserver(event string, o Observer) {
	observerRegistry[event] = append(observerRegistry[event], o)
}

func NotifyObservers(event string, msg string) {
	if observers, ok := observerRegistry[event]; ok {
		for _, o := range observers {
			o.Update(msg)
		}
	}
}
