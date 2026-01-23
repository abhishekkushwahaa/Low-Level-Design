package main

// 1. Basic Singleton (lazy initialization)
type Singleton struct{}

var instance *Singleton

func GetInstance() *Singleton {
	if instance == nil {
		instance = &Singleton{}
	}
	return instance
}
