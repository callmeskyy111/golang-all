package main

import "fmt"

// If/else on STEROIDS

//! Syntax (No 'break' needed, unlike other languages)
/*
switch expression{
case value1:
	code..
	fallthrough //Optnl.. ⬇️
case value2:
	code..
default:
	Code to be executed, if nothing matches..
}
*/

func main() {
	// Common Switch-Case
	character:= "Bruce Wayne"

	switch character{
	case "Bruce Wayne":
		fmt.Println("Owner of Wayne Enterprise 👨🏻‍💼")
	case "Batman":
		fmt.Println("Protector of Gotham City 🦇")	
	default:
		fmt.Println("Some Gotham citizen.. 👤")	
	}

	// Multiple conditions
	enemy:="Scarecrow"

	switch enemy{
	case "Bane","Scarecrow","Joker","Poison Ivy","Penguin":
		fmt.Println("Enemey from Gotham City! 🌆💀")
	case "Silver Monkey":
		fmt.Println("External enemy, from the League Of Shadows 🥷🏻⚔️")	
	default:
		fmt.Println("Enemy, within GCPD 🚓💀")

	}

	// Expressions in CASE-STATEMENTS
	num:=15

	switch{
	case num<10:
		fmt.Println("Number is less than 10")
	case num>=10 && num<20:
		fmt.Println("Number is between 10 and 19")	
	default:
		fmt.Println("Number is 20 or more")		
	}

	// Fallthrough (Optnl.)
	num2:=2
	switch{
	case num2>1:
		fmt.Println("Greater than 1")
		fallthrough // Just the next case, not all ⚠️
	case num2==2:
		fmt.Println("Number is 2")
	default:
		fmt.Println("Number isn't 2")			

	}

	// checking-type
	checkType(10)
	checkType(5.99)
	checkType("Alfred Pennywise💖")
	checkType(true)

}

// SWITCH-CASE for Type-Assertions
func checkType(x any){
	switch x.(type){
	case int:
		fmt.Println("It's an INTEGER")
	case float64:
		fmt.Println("It's a FLOAT")	
		//! fallthrough ❌ - cannot fallthrough in type switch
	case string:
		fmt.Println("It's a STRING")
	default:
		fmt.Println("UNKNOWN TYPE!")		
	}
}