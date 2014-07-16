// Program demostrates the use of using a defer to
// catch panics.
package main

import (
	"fmt"
	"runtime"
)

// main is the entry point for the program
func main() {
	if err := testPanic(); err != nil {
		fmt.Println("Test Error:", err)
	}

	if err := testPanic(); err != nil {
		fmt.Println("Test Panic:", err)
	}
}

// catchPanic handles capturing panics and reporting the problem.
func catchPanic(err *error, functionName string) {
	if r := recover(); r != nil {
		fmt.Printf("%s : PANIC Defered : %v\n", functionName, r)

		// Capture the stack trace
		buf := make([]byte, 10000)
		runtime.Stack(buf, false)

		fmt.Printf("%s : Stack Trace : %s", functionName, string(buf))

		if err != nil {
			*err = fmt.Errorf("%v", r)
		}
	}
}

// mimicError returns an error to testing the defer
func mimicError(key string) error {
	return fmt.Errorf("Mimic Error : %s", key)
}

// testError display the behavior for handling errors
func testError() (err error) {
	defer catchPanic(&err, "TestError")
	fmt.Printf("\nTestError : Start Test\n")

	err = mimicError("1")

	fmt.Printf("TestError : End Test\n")
	return err
}

// testPanic displays the behavior for handing panics
func testPanic() (err error) {
	defer catchPanic(&err, "TestPanic")
	fmt.Printf("\nTestPanic: Start Test\n")

	err = mimicError("1")

	panic("Mimic Panic")
}
