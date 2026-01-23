package main

type Singleton struct{}

var instance *Singleton

// 1. Basic Singleton (lazy initialization)
// Problem:
// a)- Not thread-safe :- a piece of code, function, or data structure is not designed to work correctly when accessed by multiple threads simultaneously
// b)- Race conditions :- A race condition is a software bug where the outcome of a program depends on the unpredictable timing or sequence of multiple processes or threads accessing shared data, leading to inconsistent or incorrect results, like two users booking the last ticket simultaneously or bank transactions miscalculating balances. It happens in concurrent systems when operations aren't properly synchronized, creating a "race" to access and modify data, often with "check-then-act" logic failing due to intervening changes. Solutions involve using synchronization tools like locks or mutexes to control access, ensuring only one thread operates at a time.

func GetInstance() *Singleton {
	if instance == nil {
		instance = &Singleton{}
	}
	return instance
}
