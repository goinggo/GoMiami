// Program demostrates the use of a defer function using
// a named return argument to capture the error.
package main

import (
	"fmt"
)

// main is the entry point for the program
func main() {
	if err := test(); err != nil {
		fmt.Println(err)
	}
}

// mimicError returns an error to testing the defer.
func mimicError(key string) error {
	return fmt.Errorf("Mimic Error : %s", key)
}

// test helps run the program logic.
func test() (err error) {
	defer func() {
		fmt.Println("Start Defer")

		if err != nil {
			fmt.Println("Err Addr:", &err)
			fmt.Println("Defer Error", err)
		}
	}()

	fmt.Println("Start Test")

	err = mimicError("1")
	fmt.Println("Err Addr:", &err)

	err = mimicError("2")
	fmt.Println("Err Addr:", &err)

	fmt.Println("End Test")
	return err
}
