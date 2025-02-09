// Where two objects can use each other but are independent.

package main

import "fmt"

type School struct {
	name string
}

type Student struct {
	name   string
	school *School // Association
}

func (s *Student) Attend() {
	fmt.Println(s.name, "is attending", s.school.name)
}

func main() {
	school := &School{name: "SCA Inter College!"}
	student := &Student{name: "Abhi", school: school}
	student.Attend()
}

// Key Takeaways = The Student can be associated with a School, but the Student and School can exist independently
