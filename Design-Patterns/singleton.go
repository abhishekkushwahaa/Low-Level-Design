package main

import (
	"fmt"
	"sync"
)

// Singleton Struct
type Singleton struct {
	value string
}

// Singleton instance (Private Field)
var instance *Singleton

// Look for thread-safety
var once sync.Once

// GetInstance returns the single instance of the singleton

func GetInstance(value string) *Singleton {
	once.Do(func() {
		instance = &Singleton{value: value}
	})
	return instance
}

// Set Value method
func (s *Singleton) SetValue(value string) {
	s.value = value
}

// Get Value Method
func (s *Singleton) GetValue() string {
	return s.value
}

func main() {
	// Get the Singleton Instance
	singleton1 := GetInstance("First Instance!")
	fmt.Println(singleton1.GetValue())

	// Get the same Singleton Instance again
	singleton2 := GetInstance("Second Instance!")
	fmt.Println(singleton2.GetValue())

	// Modify the Singleton Instance value
	singleton2.SetValue("Updated Instance!")
	fmt.Println(singleton1.GetValue())
}
