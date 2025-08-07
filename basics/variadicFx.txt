package main

import "fmt"

// variadic f(x) - can accept variable no. of args. 
// ... Ellipsis
// func funcName(param1 type1, param2 type2, param3 ...type3) returnType{
// ...fx body/code...
// }

//⚠️ variadic ...params must come last - (param paramType, vParam ...vParamType)

func main() {
	str, tripleSum:=sum("Summing 3 digits:",1,2,3,4)
	str2, doubleSum:=sum("Summing 2 digits:",2,3)

	fmt.Println(str, tripleSum)
	fmt.Println(str2, doubleSum)

	numbers:= []int{3,4,5,6,6,7,8,7,1}
	seq,tot:=sumSlice(3,numbers...)
	fmt.Printf("Sequnce no. is %d, sum is %d\n",seq,tot)
	
}

func sum(str string, nums ...int)(string,int){
 total:=0
 for _,num:=range nums{
	total+=num
 }
 return str,total
}

// With slice []
func sumSlice(sequence int, nums ...int)(int,int){
 total:=0
 for _,num:=range nums{
	total+=num
 }
 return sequence,total
}




