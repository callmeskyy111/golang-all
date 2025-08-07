package main

import "fmt"

//! In Go, a panic represents a runtime error that disrupts the normal execution flow of a program. Unlike conventional error handling, which typically returns error values, panic is used for exceptional and unrecoverable situations where the program cannot continue safely.

// panic(interface{}) or panic(any) ⚠️

func main() {
	// valid input
	process(-10)
}

func process(input int){
	defer fmt.Println("Deferred 1 ⌛")
	defer fmt.Println("Deferred 2 ⌛")
	if input<0{
		fmt.Println("This is before PANIC")
		panic("Input must be a non -ve no.🔴")
		// ❌ fmt.Println("This is before PANIC") // Unreachable code (even defer)
		
	}
	fmt.Println("Processing the input ☑️")
}