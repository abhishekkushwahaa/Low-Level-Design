# Design Patterns

## Introduction
- Design patterns are typical solutions to common problems in software design.
- Each pattern is like a blueprint that you can customize to solve a particular design problem in your code.
- Design patterns can speed up the development process by providing tested, proven development paradigms.
- Effective software design requires considering issues that may not become visible until later in the implementation.
- Reusing design patterns helps to prevent subtle issues that can cause major problems and improves code readability for coders and architects who are familiar with the patterns.

## Creational Patterns

### Singleton
- Singleton is a creational design pattern that lets you ensure that a class has only one instance, while providing a global access point to this instance.
- [Singleton Code](singleton.go)

### Factory
- Factory is a creational design pattern that provides an interface for creating objects in a superclass, but allows subclasses to alter the type of objects that will be created.
- [Factory Code](factory.go)

### Abstract Factory
- Abstract Factory is a creational design pattern that lets you produce families of related objects without specifying their concrete classes.
- [Abstract Factory Code](abstract-factory.go)

### Prototype
- Prototype is a creational design pattern that lets you copy existing objects without making your code dependent on their classes.
- [Prototype Code](prototype.go)

### Builder
- Builder is a creational design pattern that lets you construct complex objects step by step. The pattern allows you to produce different types and representations of an object using the same construction code.
- [Builder Code](builder.go)

## Behavioral Patterns

### Strategy
- Strategy is a behavioral design pattern that lets you define a family of algorithms, put each of them into a separate class, and make their objects interchangeable.
- [Strategy Code](strategy.go)

### Observer
- Observer is a behavioral design pattern that lets you define a subscription mechanism to notify multiple objects about any events that happen to the object they’re observing.
- [Observer Code](observer.go)

## Structural Patterns

### Decorator
- Decorator is a structural design pattern that lets you attach new behaviors to objects by placing these objects inside special wrapper objects that contain the new behaviors.
- [Decorator Code](decorator.go)
