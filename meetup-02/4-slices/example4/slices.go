// This program demostrates how slices grow internally

package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// main is the entry point for the program
func main() {
	// Create an empty slice
	data := []string{}

	for record := 0; record < 1050; record++ {
		// Append data to the slice
		data = append(data, fmt.Sprintf("Rec: %d", record))

		// Display the details of the slice on these boundaries
		if record < 10 || record == 256 || record == 512 || record == 1024 {
			sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&data)))

			fmt.Printf("Index[%d] Len[%d] Cap[%d]\n",
				record,
				sliceHeader.Len,
				sliceHeader.Cap)
		}
	}
}

// Within the first 1k of elements, capacity is doubled.
// Then capacity is grown by a factor of 1.25 or 25%
