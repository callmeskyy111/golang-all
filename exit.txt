package main

import (
	"fmt"
	"os"
)

// Exit causes the current program to exit with the given status code. Conventionally, code zero indicates success, non-zero an error. The program terminates immediately; deferred functions are not run.

// For portability, the status code should be in the range [0, 125].

// Terminate immediately (ignoring all defers, panics, recoveries, etc.)

func main() {
	defer fmt.Println("Deferred satement ⌛")
	fmt.Println("Starting the main f(x).. ✅")

	// Exit with status code of 1
	os.Exit(1)

	// This will never be executed
	fmt.Println("End of main f(x).. ☑️")
}