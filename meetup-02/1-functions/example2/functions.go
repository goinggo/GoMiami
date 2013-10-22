// This program demostrates a shadowing error that can occur
// with naked returns

package main

import (
	"fmt"
)

// main is the entry point for the program
func main() {
	id, err := ReturnId_Shadowing()

	if err != nil {
		fmt.Printf("ERROR: %s", err)
		return
	}

	fmt.Printf("Id: %d\n", id)
}

// ReturnId_Shadowing provides an example of returning two values
// using a naked return with shadowing error
func ReturnId_Shadowing() (id int, err error) {
	id = 30

	if id == 30 {
		err := fmt.Errorf("Invalid Id\n")
		return
	}

	return
}
