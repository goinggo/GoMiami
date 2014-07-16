// Program demostrates how to return multiple values from
// a function. It also shows how to use named return arguments
// and naked returns
package main

import (
	"fmt"
)

// main is the entry point for the program.
func main() {
	id, err := ReturnID()
	if err != nil {
		fmt.Printf("ERROR: %s", err)
		return
	}

	fmt.Printf("Id: %d\n", id)

	id2, err := ReturnIDNamed()
	if err != nil {
		fmt.Printf("ERROR: %s", err)
		return
	}

	fmt.Printf("Id: %d\n", id2)

	id3, err := ReturnIDNaked()
	if err != nil {
		fmt.Printf("ERROR: %s", err)
		return
	}

	fmt.Printf("Id: %d\n", id3)
}

// ReturnID provides an example of returning two values.
func ReturnID() (int, error) {
	id := 10
	return id, nil
}

// ReturnIDNamed provides an example of returning two values
// using named return arguments.
func ReturnIDNamed() (id int, err error) {
	id = 20
	return id, err
}

// ReturnIDNaked provides an example of returning two values
// using a naked return.
func ReturnIDNaked() (id int, err error) {
	id = 30
	return
}
