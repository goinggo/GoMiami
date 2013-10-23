// This program demostrates passing by value. In this case we
// are passing a pointer

package main

import (
	"fmt"
)

func main() {
	value := 10
	ByValue(&value)

	fmt.Printf("Value %d\n", value)
}

func ByValue(value *int) {
	*value = 20
}
