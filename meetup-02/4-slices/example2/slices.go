// Program demostrates how to manipulate slices and
// create new slices from slices.
package main

import (
	"fmt"
	"unsafe"
)

// main is the entry point for the program.
func main() {
	orgSlice := make([]string, 5, 8)
	orgSlice[0] = "Apple"
	orgSlice[1] = "Orange"
	orgSlice[2] = "Banana"
	orgSlice[3] = "Grape"
	orgSlice[4] = "Plum"

	fmt.Println("Display Original Slice")
	InspectSlice(orgSlice)

	fmt.Println("From Org Slice: Slice 2 elements between indexes 2 and 3")
	slice2 := orgSlice[2:4]
	InspectSlice(slice2)

	fmt.Println("From Slice 2: Slice index position 1 up to the capacity of slice2")
	slice3 := slice2[1:cap(slice2)]
	InspectSlice(slice3)

	fmt.Println("From Slice 3: Change the value of index 0 in slice 3")
	slice3[0] = "CHANGED"
	InspectSlice(slice3)

	fmt.Println("Display slice 2")
	InspectSlice(slice2)
}

// InspectSlice display the memory and values of the slices data structure.
func InspectSlice(slice []string) {

	// Capture the address to the slice structure.
	address := unsafe.Pointer(&slice)

	// Capture the address where the length and cap size is stored.
	lenAddr := uintptr(address) + uintptr(8)
	capAddr := uintptr(address) + uintptr(16)

	// Create pointers to the length and cap size.
	lenPtr := (*int)(unsafe.Pointer(lenAddr))
	capPtr := (*int)(unsafe.Pointer(capAddr))

	// Create a pointer to the underlying array.
	addPtr := (*[8]string)(unsafe.Pointer(*(*uintptr)(address)))

	fmt.Printf("Slice Addr[%p] Len Addr[0x%x] Cap Addr[0x%x]\n",
		address,
		lenAddr,
		capAddr)

	fmt.Printf("Slice Length[%d] Cap[%d]\n",
		*lenPtr,
		*capPtr)

	for index := 0; index < *lenPtr; index++ {
		fmt.Printf("[%d] %p %s\n",
			index,
			&(*addPtr)[index],
			(*addPtr)[index])
	}

	fmt.Printf("\n\n")
}
