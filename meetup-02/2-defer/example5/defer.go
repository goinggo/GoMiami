// This program demostrates how the defer function is
// evaluated and values that are passed in are taken
// at the time of the evaluation

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
}

// ReturnId returns an id and evaluates if the id
// is valid
func ReturnId() (id int, err error) {
	defer func(id int) {
		if id == 10 {
			err = fmt.Errorf("Invalid Id\n")
		}

		fmt.Printf("Value of Id : %d\n", id)
	}(id)

	id = 10

	return id, err
}
