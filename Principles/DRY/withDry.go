// DRY = Don't Repeat Yourself
package main

import "fmt"

type User struct {
	name  string
	email string
}

// Reusable function to greet a user
func greet(user User) {
	fmt.Printf("Hello %s\n", user.name)
}

func main() {
	// Create users
	user1 := User{name: "Abhishek!", email: "abhishek@gmail.com"}
	user2 := User{name: "Balji!", email: "balji@gmail.com"}

	// Greet users using the reusable function
	greet(user1)
	greet(user2)
}
