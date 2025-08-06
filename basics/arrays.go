package main

import "fmt"

// syntax: var arrayName [size] elementType
// arrs[] := Ordered collections of data

func main() {
	
	var nums[5] int
	fmt.Println(nums)

	nums[0] = 1
	nums[len(nums)-1] = 9
	fmt.Println(nums)

	characters:= [4]string{"Batman 🦇","Alfred 💁🏻‍♂️","Harvey 🎭","Gorodon 👮🏻‍♂️"}
	fmt.Println(characters)

	// A change in the copy of the array doesn't affect the original one

	originalArr:= [3]int{1,2,3}
	copiedArr:= originalArr
	copiedArr[1]= 100

	
	fmt.Println("Original Arr :",originalArr)
	fmt.Println("Copied Arr :",copiedArr)

	fmt.Println("--------------------------")

	// Iterating over arrays
	// for i:=0; i<len(nums);i++{
	// 	fmt.Println("Idx:",i," Vals:",nums[i])
	// }

	//! Better way- range
	fmt.Println("--------------------------")
	for i:=range len(nums){
		fmt.Println("Idx:",i," Vals:",nums[i])
	}

	fmt.Println("--------------------------")

	for i,v:=range nums{
		fmt.Println("Idx:",i," Vals:",v)
	}

	fmt.Println("--------------------------")
	for _,v:=range nums{
		fmt.Println("Vals:",v)
	}

	fmt.Println("--------------------------")
	for i,_:=range nums{
		fmt.Println("Idx:",i)
	}

	fmt.Println("--------------------------")


	fmt.Println("The length of num[ ] is:",len(nums))

	arr1:= [3]int{1,2,3}
	arr2:= [3]int{1,2,3}
	fmt.Println(arr1==arr2) // true

	// Multidimensional Array
	fmt.Println("--------------------------")

	var matrix[3][3]int = [3][3]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}

	fmt.Println(matrix) // [[1 2 3] [4 5 6] [7 8 9]]
	

	// Using ptrs* to confirm change 🧠🌟
	fmt.Println("--------------------------")

	originalArr1:= [3]int{11,22,33}
	copiedArr1:= &originalArr1
	copiedArr1[2]= 1000
	
	fmt.Println("Original Arr1 :",originalArr1)
	fmt.Println("Copied Arr1 :",*copiedArr1)

}