package main

const PI = 3.14 // untyped constants (inference)
const GRAVITY = 9.81

func main() {
	const days int = 7 // typed constants
	const (
		monday        = 1
		tuesday       = 2
		wednesday     = 3
		thursday  int = 4
	)
}