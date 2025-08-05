package main

import "fmt"

//! SYNTAX
//if condition {
// block of code
//}else if condition{
// another block of code
//}else{
// last code-block
//}



func main() {

	age:=29
	temperature:=25

	//if
	if age>=18{
		fmt.Println("You're an adult!")
	}

	//if-else
	if temperature>=30{
		fmt.Println("It's hot outside")
	} else {
		fmt.Println("It's cool outside")
	}

	//if->else if->else
	score:= 85
	if score >=90 {
		fmt.Println("Grade A")
	}else if score>=80{
		fmt.Println("Grade B")
	}else if score>=70{
		fmt.Println("Grade C")	
	}else{
		fmt.Println("Grade D")
	}

	// Nested conditionals
	num:=18
	if num%2==0{
		if num % 3 == 0{
			fmt.Println("Number is divisible by both 2 ✅ & 3 ☑️")
		}else{
			fmt.Println("Number is divisible by 2✅, but not 3❌")
		}
	}else{
		fmt.Println("Number is not divisible by 2 ❌")
	}

	// Also- 
	// || OR
	// && AND
	
}