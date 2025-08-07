package main

import (
	"fmt"
	"slices"
)

// More dynamic than plain arrays[]
// Do not have fixed lengths
// More commonly used

// syntax: var sliceName[]ElenmentType


func main() {
	// var nums []int
	// var nums1 = []int{1,2,3}

	// nums2:= []int{9,8,7}

	// Using make()
	// makeSlice := make([]int,5)
	a:= [5]int{1,2,3,4,5}

	slice1:= a[1:4]

	fmt.Println("Slice1:",slice1)

	slice1 = append(slice1, 6,7)
	fmt.Println("Appended Slice1:",slice1)

	sliceCopy:= make([]int,len(slice1))
	copy(sliceCopy,slice1)
	fmt.Println("Slice-Copy:",sliceCopy)

	// Now, nil slices
	//var nilSlice []int

	for i,v:=range slice1{
		fmt.Println(i,v)
	}
	fmt.Println("Element at idx 3 of slice1:", slice1[3])
	slice1[3] = 50
		fmt.Println("Element at idx 3 of slice1:", slice1[3])
	
	if slices.Equal(slice1,sliceCopy){
		fmt.Println("slice1 is equal to sliceCopy ✅")
	}else{
		fmt.Println("slice1 is NOT equal to sliceCopy ❌")
	}	

	// Multidimensional Slices[][]
	twoD := make([][]int, 3) // 1️⃣

	for i := range twoD { // 2️⃣
	innerLen := i + 1 // 3️⃣
	twoD[i] = make([]int, innerLen) // 4️⃣

	for j := range twoD[i] {
		twoD[i][j] = i + j // 5️⃣
		}
	}

	fmt.Println(twoD) // 6️⃣

	// slice[low:high] // slice op.
	slice2:= slice1[2:4]
	fmt.Println(slice2)
	fmt.Println("Length of slice2: ",len(slice2))
	fmt.Println("Capacity of slice2: ",cap(slice2))

}