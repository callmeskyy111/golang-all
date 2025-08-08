package main

import (
	"fmt"
	"math/rand"
	"time"
)

// for loop As while loop


func main() {

	//i:=1
	// sum:=4000
	// for {
	// 	sum+=1
	// 	fmt.Println("Sum: ",sum)
	// 	//i++
	// 	if sum>=10000{
	// 		fmt.Println("Reached limit! 10,000.. 🔴")
	// 		break
	// 	}

	// }

	//! Guessing Game 🎮❓
	src:= rand.NewSource(time.Now().UnixNano())
	random:= rand.New(src)

	// Generate a random number between 1 and 100
	target:= random.Intn(100)+1

	// Welcome message(s)
	fmt.Println("🤖: Welcome to the GUESSING GAME!")
	fmt.Println("🤖: I have chosen a number between 1-100")
	fmt.Println("🤖: Can you guess the number?")

	var guess int

	for {
		fmt.Println("🤖: Enter your guess:")
		fmt.Scanln(&guess)

		// Check if the guess is correct!
		if guess == target{
			fmt.Println("Congrats! You guessed the correct number 🎉")
			break
		} else if guess < target{
			fmt.Println("TOO LOW! Try again.. 🔴")
			//break
		} else{
			fmt.Println("TOO HIGH! Try again.. 🔴")
			//break
		}
	}



}