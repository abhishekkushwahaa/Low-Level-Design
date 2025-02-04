// DRY = Don't Repeat Yourself
package main

import "fmt"

type User struct {
	name  string
	email string
}

func main() {
	// Create users
	user1 := User{name: "Abhishek!", email: "abhsihek@gmail.com"}
	user2 := User{name: "Balji!", email: "balji@gmail.com"}

	// Repeated
	fmt.Printf("Hello %s\n", user1.name)
	fmt.Printf("Hello %s\n", user2.name)
}
