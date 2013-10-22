// This program demostrates how to return multiple values from
// a function. It also shows how to use named return arguments
// and naked returns

package main

import (
	"fmt"
)

// main is the entry point for the program
func main() {
	id, err := ReturnId()

	if err != nil {
		fmt.Printf("ERROR: %s", err)
		return
	}

	fmt.Printf("Id: %d\n", id)

	id2, _ := ReturnId_Named()
	fmt.Printf("Id: %d\n", id2)

	id3, _ := ReturnId_Naked()
	fmt.Printf("Id: %d\n", id3)
}

// ReturnId provides an example of returning two values
func ReturnId() (int, error) {
	id := 10
	return id, nil
}

// ReturnId_Named provides an example of returning two values
// using named return arguments
func ReturnId_Named() (id int, err error) {
	id = 20
	return id, err
}

// ReturnId_Naked provides an example of returning two values
// using a naked return
func ReturnId_Naked() (id int, err error) {
	id = 30
	return
}
