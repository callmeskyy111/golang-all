package main

import "fmt"

//! func recover() any
// The recover built-in function allows a program to manage behavior of a panicking goroutine. Executing a call to recover inside a deferred function (but not any function called by it) stops the panicking sequence by restoring normal execution and retrieves the error value passed to the call of panic. If recover is called outside the deferred function it will not stop a panicking sequence. In this case, or when the goroutine is not panicking, recover returns nil.
// Only useful inside deferred f(x)
// Prior to Go 1.21, recover would also return nil if panic is called with a nil argument

func main() {
	process()
	fmt.Println("Returned from process!")
}

func process(){
	defer func ()  {
		 r:=recover();
		 if r!=nil{
			fmt.Println("Recovered ✅",r)
		 }
	}()
	fmt.Println("Start Process...")
	panic("\nSomething went wrong 🔴")
	//fmt.Println("Recovered from PROCESS!")
}