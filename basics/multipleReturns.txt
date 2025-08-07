package main

import (
	"errors"
	"fmt"
)

// func funcName(param1 type1, param2 type2,....)(returnType1, returnType2,....){
// code block...
// return returnValue1, returnValue2...
// }

func main() {
	q,r:=divide(10,3)
	fmt.Printf("Quotient: %d, Reminder: %d\n",q,r)

	res,err:=compare(3,3)
	if err!=nil{
		fmt.Println("🔴 Error:",err)
	}else{
		fmt.Println(res)
	}

}

func divide(a,b int)(int,int){
	quotient:=a/b
	remainder:=a%b
	return quotient, remainder
}

func compare(a, b int)(string, error){
	if a>b {
		return "a is greater than b",nil
	}else if b>a{
		return "b is greater than a",nil
	}else{
		return "",errors.New("Unable to compare!")
	}
	
}