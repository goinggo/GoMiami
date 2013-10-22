// This program demostrates an object oriented program. It demostrates
// the use of composition and interfaces

package main

import (
	"fmt"
)

// Animal contains based related properties
type Animal struct {
	Name string
	mean bool
}

// AnimalSounder is an interface to be implemented by Animals
type AnimalSounder interface {
	MakeNoise()
}

// Dog represents a dog animal and its specific properties
type Dog struct {
	Animal
	BarkStrength int
}

// Cat represents a cat animal and its specific properties
type Cat struct {
	Basics       Animal
	MeowStrength int
}

// main is the entry point for the program
func main() {

	// Use a composite literal
	myDog := &Dog{
		Animal{
			"Rover", // Name
			false,   // mean
		},
		2, // BarkStrength
	}

	// Use a composite literal
	myCat := &Cat{
		Basics: Animal{
			Name: "Julius",
			mean: true,
		},
		MeowStrength: 3,
	}

	MakeSomeNoise(myDog)
	MakeSomeNoise(myCat)
}

// PerformNoise performs the actual work of making an animal sound
func (animal *Animal) PerformNoise(strength int, sound string) {
	if animal.mean == true {
		strength = strength * 5
	}

	for voice := 0; voice < strength; voice++ {
		fmt.Printf("%s ", sound)
	}

	fmt.Printf("\n")
}

// MakeNoise performs the barking for a Dog
// Implements the AnimalSounder interface
func (dog *Dog) MakeNoise() {
	dog.PerformNoise(dog.BarkStrength, "BARK")
}

// MakeNoise perform the barking for a Cat
// Implements the AnimalSounder interface
func (cat *Cat) MakeNoise() {
	cat.Basics.PerformNoise(cat.MeowStrength, "MEOW")
}

// MakeSomeNoise uses the interface to to cause each respective animal
// make noise
func MakeSomeNoise(animalSounder AnimalSounder) {
	animalSounder.MakeNoise()
}
