// This program demostrates the use of a defer function using
// a named return argument to capture the error

package main

import (
	"errors"
	"fmt"
)

// main is the entry point for the program
func main() {
	err := Test()

	if err != nil {
		fmt.Printf("%s\n", err)
	}
}

// MimicError returns an error to testing the defer
func MimicError(key string) error {
	return errors.New(fmt.Sprintf("Mimic Error : %s", key))
}

// Test helps run the program logic
func Test() (err error) {
	defer func() {
		fmt.Printf("Start Defer\n")

		if err != nil {
			fmt.Printf("Err Addr: %v\n", &err)
			fmt.Printf("Defer Error : %v\n", err)
		}
	}()

	fmt.Printf("Start Test\n")

	err = MimicError("1")
	fmt.Printf("Err Addr: %v\n", &err)

	err = MimicError("2")
	fmt.Printf("Err Addr: %v\n", &err)

	fmt.Printf("End Test\n")

	return err
}
