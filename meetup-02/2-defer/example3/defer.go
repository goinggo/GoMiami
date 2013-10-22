// This program demostrates the use of using two defer functions

package main

import (
	"errors"
	"fmt"
)

// main is the entry point for the program
func main() {
	Test()
}

// MimicError returns an error to testing the defer
func MimicError(key string) error {
	return errors.New(fmt.Sprintf("Mimic Error : %s", key))
}

// Test helps run the program logic
func Test() (err error) {
	defer func() {
		fmt.Printf("Start Panic Defer\n")

		if r := recover(); r != nil {
			fmt.Printf("Defer Panic : %v\n", r)
		}
	}()

	defer func() {
		fmt.Printf("Start Defer\n")

		if err != nil {
			fmt.Printf("Defer Error : %v\n", err)
		}
	}()

	fmt.Printf("Start Test\n")

	err = MimicError("1")

	panic("Mimic Panic")

	fmt.Printf("End Test\n")

	return err
}
