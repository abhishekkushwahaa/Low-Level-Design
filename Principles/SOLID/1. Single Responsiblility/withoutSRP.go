package main

import "fmt"

type UserService struct{}

func (u *UserService) User(name, email string) {
	// Responsibility 1: Create User
	fmt.Println("User Created:", name)
	// Responsibility 2: Send Email
	fmt.Println("Sending Email to:", email)
}

func main() {
	user := &UserService{}
	user.User("Abhishek G", "abhishek@gmail.com")
}
