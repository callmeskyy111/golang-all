package main

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	//! PascalCase
	// Structs, interfaces, enums
	// Ex.- CalculateArea, UserInfo{}, NewHttpRequest etc.

	//! snake_case
	// Vars, contsants, files
	// Ex.- user_id, first_name, http_request etc.

	//! UPPERCASE
	// constants (only!)
	// PI, FINALPRICE etc.

	//! mixedCase
	// External libraries, other languages
	//Ex.- javaScript, isValid etc.

	const MAX_RETRIES = 5

	employeeId:=1001
	fmt.Println("EmployeeId:",employeeId)
}