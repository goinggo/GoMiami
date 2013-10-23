// This program demostrates the use of a defer function

package main

import (
	"fmt"
)

// main is the entry point for the program
func main() {
	Test()
}

// MimicError returns an error to testing the defer
func MimicError(key string) error {
	return fmt.Errorf("Mimic Error : %s", key)
}

// Test helps run the program logic
func Test() {
	fmt.Printf("Start Test\n")

	// Short variable declaration
	err := MimicError("1")
	fmt.Printf("Err Addr: %v\n", &err)

	defer func() {
		fmt.Printf("Start Defer\n")

		if err != nil {
			fmt.Printf("Err Addr: %v\n", &err)
			fmt.Printf("Defer Error : %v\n", err)
		}
	}()

	err = MimicError("2")

	fmt.Printf("End Test\n")
}
