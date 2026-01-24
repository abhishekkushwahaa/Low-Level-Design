package main

import "sync"

type Singleton struct{}

var instance *Singleton

// 1. Basic Singleton (lazy initialization)
// Benefits: Created only when needed (better memory usage)
// Problems:
// a)- Not thread-safe :- a piece of code, function, or data structure is not designed to work correctly when accessed by multiple threads simultaneously
// b)- Race conditions :- A race condition is a software bug where the outcome of a program depends on the unpredictable timing or sequence of multiple processes or threads accessing shared data, leading to inconsistent or incorrect results, like two users booking the last ticket simultaneously or bank transactions miscalculating balances. It happens in concurrent systems when operations aren't properly synchronized, creating a "race" to access and modify data, often with "check-then-act" logic failing due to intervening changes. Solutions involve using synchronization tools like locks or mutexes to control access, ensuring only one thread operates at a time.

func GetInstance() *Singleton {
	if instance == nil {
		instance = &Singleton{}
	}
	return instance
}

// 2. Eager Initialization
// Benefits: Simple and safe
// Problems: Not lazy - allocated whether used or not
func GetEagerInstance() *Singleton {
	return instance
}

// 3. Thread-safe Singleton
// Benefits: Thread-safe
// Problems: Slow (locks every call)
var mu sync.Mutex

func GetInstanceThreadSafe() *Singleton {
	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		instance = &Singleton{}
	}
	return instance
}

// 4. Double-Checked Locking
// Benefits: a)- Faster than always locking b)- Safe in Go due to memory model (mostly)
// Problems: a)- In some languages, unsafe without memory barriers b)- More complex logic
func GetInstanceSafeDoubleCheck() *Singleton {
	if instance == nil {
		mu.Lock()
		defer mu.Unlock()
		if instance == nil {
			instance = &Singleton{}
		}
	}
	return instance
}

// 5. sync.Once - Best in Go
// Benefits: a)- Lazy b)- Thread-safe c)- No locking overhead after initial call
var once sync.Once

func GetInstanceOnce() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}

// 6. Closure-Based Singleton: Encapsulates instance inside returned function.
// Benefits: a)- True encapsulation b)- No global variable exposed
var GetClosureInstance = func() func() *Singleton {
	var instance *Singleton // inside closure

	return func() *Singleton {
		if instance == nil {
			instance = &Singleton{}
		}
		return instance
	}
}
