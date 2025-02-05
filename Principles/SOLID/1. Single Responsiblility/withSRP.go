package main

import "fmt"

type UserCreate struct{}

func (*UserCreate) CreateUser(name, email string) {
	fmt.Println("User Created:", name, "&", email)
}

type EmailService struct{}

func (e *EmailService) SendEmail(email string) {
	fmt.Println("Sending Email to:", email)
}

func main() {
	userService := &UserCreate{}
	emailService := &EmailService{}

	userService.CreateUser("Abhishek G", "abhishek@gmail.com")
	emailService.SendEmail("abhishek@gmail.com")
}
