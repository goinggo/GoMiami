// Program demostrates the use of using two defer functions.
package main

import (
	"fmt"
)

// main is the entry point for the program.
func main() {
	test()
}

// mimicError returns an error to testing the defer.
func mimicError(key string) error {
	return fmt.Errorf("Mimic Error : %s", key)
}

// Test helps run the program logic.
func test() (err error) {
	defer func() {
		fmt.Println("Start Panic Defer")

		if r := recover(); r != nil {
			fmt.Println("Defer Panic", r)
		}
	}()

	defer func() {
		fmt.Println("Start Defer")

		if err != nil {
			fmt.Println("Defer Error", err)
		}
	}()

	fmt.Println("Start Test")

	err = mimicError("1")

	panic("Mimic Panic")
}
