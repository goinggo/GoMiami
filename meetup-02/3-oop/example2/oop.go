// This program demostrates an object oriented program that uses an
// interface as a type

package main

import (
	"fmt"
)

// HornSounder is an interface for sounding a horn
type HornSounder interface {
	SoundHorn()
}

// Vehicles is a two element array of objects that can sound a horn
type Vehicles [2]HornSounder

// Car represents a Vechicle with a sound
type Car struct {
	Sound string
}

// Bike represents a Vechicle with a sound
type Bike struct {
	Sound string
}

// main is the entry point for the program
func main() {

	// Use a composite literal to create an object of type Vechicles
	vehicles := &Vehicles{}

	// Add a Car
	vehicles[0] = &Car{
		Sound: "BEEP",
	}

	// Add a Bike
	vehicles[1] = &Bike{
		Sound: "RING",
	}

	// Sound the horn for all listed vehicles
	for _, hornSounder := range vehicles {
		hornSounder.SoundHorn()
	}
}

// SoundHorn implements the interface for the Car type
func (car *Car) SoundHorn() {
	fmt.Println(car.Sound)
}

// SoundHorn implements the interface for the Bike type
func (bike *Bike) SoundHorn() {
	fmt.Println(bike.Sound)
}
