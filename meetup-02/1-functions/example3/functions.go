// Program demostrates a use case where using
// named return arguments don't make sense.
package main

import (
	"fmt"
)

// main is the entry point for the program.
func main() {
	ans := AddNumbers(10, 12)
	fmt.Printf("Answer: %d\n", ans)
}

// AddNumbers add the specified numbers and returns the result.
func AddNumbers(a int, b int) (result int) {
	return a + b
}
