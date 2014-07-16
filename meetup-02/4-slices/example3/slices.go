// Program demostrates how to append data to slices and
// grow their capacity.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Record is a sample structure for the example.
type Record struct {
	ID    int
	Name  string
	Color string
}

// main is the entry point for the program
func main() {
	// Let's keep things unknown
	random := rand.New(rand.NewSource(time.Now().Unix()))

	// Create a slice with a large length and capacity pretending we retrieved data
	// from a database
	data := make([]Record, 1000)

	// Create the data set
	for record := 0; record < 1000; record++ {
		pick := random.Intn(10)
		color := "Red"

		if pick == 2 {
			color = "Blue"
		}

		data[record] = Record{
			ID:    record,
			Name:  fmt.Sprintf("Rec: %d", record),
			Color: color,
		}
	}

	// Create two empty slices and split the records by color
	var red []Record
	var blue []Record

	for _, record := range data {
		if record.Color == "Red" {
			red = append(red, record)
		} else {
			blue = append(blue, record)
		}
	}

	// Display the counts
	fmt.Printf("Red[%d] Blue[%d]\n", len(red), len(blue))
}
