// KISS = Keep it Simple, Stupid
package main

import "fmt"

func checkNumber(n int) {
	if n > 0 {
		fmt.Println("Positive Number!")
	} else {
		if n < 0 {
			fmt.Println("Negative Number!")
		} else {
			fmt.Println("Zero!")
		}
	}
}

func main() {
	checkNumber(5)
}
