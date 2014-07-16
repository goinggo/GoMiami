// Program demostrates how the defer function is
// evaluated and values that are passed in are taken
// at the time of the evaluation.
package main

import (
	"errors"
	"fmt"
)

// main is the entry point for the program.
func main() {
	id, err := returnID()
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	fmt.Println("Id:", id)
}

// returnID returns an id and evaluates if the id is valid.
func returnID() (id int, err error) {
	defer func(id int) {
		if id == 10 {
			err = errors.New("Invalid Id")
		}

		fmt.Println("Value of Id:", id)
	}(id)

	id = 10
	return id, err
}
