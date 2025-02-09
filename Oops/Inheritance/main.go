// A new class/struct to inherit the properties and behaviors of an existing class.

package main

import "fmt"

type Animal struct {
	name string
}

func (a *Animal) Speak() {
	fmt.Println("Animal speaks")
}

type Dog struct {
	Animal
}

func (d *Dog) Speak() {
	fmt.Println(d.name + "barks")
}

func main() {
	animal := Animal{name: "Generic Animal"}
	dog := Dog{Animal: Animal{name: "German Shephard "}}
	animal.Speak()
	dog.Speak()
}
