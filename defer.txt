package main

import "fmt"

//! In Golang, the defer keyword is used to schedule a function call to be executed just before the surrounding function returns. This ensures that certain cleanup operations, like closing files or releasing resources, are performed regardless of how the function exits (e.g., through a normal return, an error, or a panic). 🌟

//! Multiple deferred statements? - LIFO (3️⃣.. 2️⃣.. 1️⃣..)


func main() {
	process(2)
}

func process(i int){
	defer fmt.Println("⌛ Deferred value of i:",i)
	defer fmt.Println("⌛ 1st deferred statement execution 1️⃣")
	defer fmt.Println("⌛ 2nd deferred statement execution 2️⃣")
	defer fmt.Println("⌛ 3rd deferred statement execution 3️⃣")

	i++

	fmt.Println("Normal statement execution ☑️")
	fmt.Println("Normal value of i:",i,"☑️")
}