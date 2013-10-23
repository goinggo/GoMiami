// This program fixes the bug surrounding the problems spawing the
// go routine

package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

// Shutdown is a package level variable to flag
// a shutdown should take place early
var Shutdown bool = false

// Kill the program after the timeout has been reached
var TimeoutSeconds int = 10

// main is the entry point for the program
func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	complete := make(chan struct{})
	go LaunchProcessor(complete)

	for {
		select {
		case whatSig := <-sigChan:
			Shutdown = true
			continue

		case <-time.After(time.Duration(TimeoutSeconds) * time.Second):
			fmt.Printf("******> TIMEOUT\n")
			os.Exit(1)

		case <-complete:
			return
		}
	}
}

// LaunchProcessor is a go routine that is spawned to
// simulate work
func LaunchProcessor(complete chan struct{}) {
	defer func() {
		close(complete)
	}()

	fmt.Printf("Start Work\n")

	for count := 0; count < 5; count++ {
		fmt.Printf("Doing Work\n")
		time.Sleep(1 * time.Second)

		if Shutdown == true {
			fmt.Printf("Kill Early\n")
			return
		}
	}

	fmt.Printf("End Work\n")
}
