package main

import "fmt"

//! range creates a copy, modifies the copied value

// works with: arrays,slices,maps,channels.. etc.

func main() {
	batman := "I am vengeance"

	for i, v := range batman {
		fmt.Println(i,v) // idx, unicode-vals.
		fmt.Printf("Idx: %d, Rune: %c\n",i,v) // idx, runes/chars
	}
	// Also.. _,v , i,_ works too!
	// _ helps in preventing memory leaks!
}