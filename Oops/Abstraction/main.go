// The concept of hiding the complex implementation details and only exposing the necessary functionalities.

package main

import "fmt"

type Device interface {
	TurnOn()
	TurnOff()
}

type TV struct{}

func (t *TV) TurnOn() {
	fmt.Println("Turn on the TV")
}

func (t *TV) TurnOff() {
	fmt.Println("Turn off the TV")
}

type Remote struct {
	device Device
}

func (r *Remote) Control() {
	r.device.TurnOn() // Abstract intraction
}

func main() {
	tv := &TV{}
	remote := &Remote{device: tv}
	remote.Control() // User controls TV without knowing its internal mechanism
}

// Key takeaways = Abstraction allows you to interact with high-level functionalities without worrying about the details
