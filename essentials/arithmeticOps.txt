package main

import (
	"fmt"
	"math"
)

func main() {
	// Var. declaration
	var a, b int = 10, 3

	var result int = a + b
	fmt.Println("Addition:",result)

	result = a - b
	fmt.Println("Subtraction:",result)

	result = a*b
	fmt.Println("Product:",result)

	result  = a/b
	fmt.Println("Division:",result)

	result = a%b
	fmt.Println("Modulus/Remainder:",result)

	const p float64 = 22/7.0
	fmt.Println(p)

	// Be mindful of OVERFLOW and UNDERFLOW
	var maxInt int64 = 9223372036854775807
	fmt.Println(maxInt)

	// Overflow with signed int.
	maxInt +=1
	fmt.Println(maxInt) // -9223372036854775808

	// Overflow with unsigned int.
	var uMaxInt uint = 18446744073709551615
	uMaxInt+=1
	fmt.Println(uMaxInt) // 0

	// Underflow with floating point numbers
	var smallFloat float64 = 1.0e-323
	smallFloat/=math.MaxFloat64
	fmt.Println(smallFloat) // 0
}

//! Output:
// Addition: 13
// Subtraction: 7
// Product: 30
// Division: 3
// Modulus/Remainder: 1
// 3.142857142857143
// 9223372036854775807
// -9223372036854775808
// 0
// 0