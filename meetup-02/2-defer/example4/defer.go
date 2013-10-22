// This program demostrates the use of using a defer to
// catch panics

package main

import (
	"errors"
	"fmt"
	"runtime"
)

// main is the entry point for the program
func main() {
	var err error

	err = TestError()

	if err != nil {
		fmt.Printf("Test Error: %v\n", err)
	}

	err = TestPanic()

	if err != nil {
		fmt.Printf("Test Panic: %v\n", err)
	}
}

// catchPanic handles capturing panics and reporting the problem
func catchPanic(err *error, functionName string) {
	if r := recover(); r != nil {
		fmt.Printf("%s : PANIC Defered : %v\n", functionName, r)

		// Capture the stack trace
		buf := make([]byte, 10000)
		runtime.Stack(buf, false)

		fmt.Printf("%s : Stack Trace : %s", functionName, string(buf))

		if err != nil {

			*err = errors.New(fmt.Sprintf("%v", r))
		}
	}
}

// MimicError returns an error to testing the defer
func MimicError(key string) error {
	return errors.New(fmt.Sprintf("Mimic Error : %s", key))
}

// TestError display the behavior for handling errors
func TestError() (err error) {
	defer catchPanic(&err, "TestError")

	fmt.Printf("\nTestError : Start Test\n")

	err = MimicError("1")

	fmt.Printf("TestError : End Test\n")

	return err
}

// TestPanic displays the behavior for handing panics
func TestPanic() (err error) {
	defer catchPanic(&err, "TestPanic")

	fmt.Printf("\nTestPanic: Start Test\n")

	err = MimicError("1")

	panic("Mimic Panic")

	fmt.Printf("TestPanic: End Test\n")

	return err
}
