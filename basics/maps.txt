package main

import (
	"fmt"
	"maps"
)

// Syntax: var m map[keyType]valueType
// Syntax2: mapVariable = make(map[keyType]valueType)
// Syntax3: mapVariable = map[leyType]valueType{ key1: value1, ..., keyN: valueN} // map-literal

func main() {
	myMap:=make(map[string]int)
	fmt.Println(myMap)

	myMap["age"] = 29
	fmt.Println(myMap)

	myMap["std-code"] = 91
	fmt.Println(myMap)

	myMap["pin-code"] = 700079
	fmt.Println(myMap)

	fmt.Println(myMap["age"])

	// Non-existent key : default value (0/""/nil)
	myMap["pin-code"] = 700028
	fmt.Println(myMap)

	// Deletion
	delete(myMap,"pin-code")
	fmt.Println(myMap)

	// Clearing the map completely
	//clear(myMap)


	//fmt.Println("After CLEAR:",myMap) // map[]

	val, unknownVal:=myMap["age"]
	fmt.Println(val) // 29
	fmt.Println(unknownVal) // true

	// Another example
	myMap2:=map[string]int{"a":1,"b":2}
	myMap3:=map[string]int{"a":1,"b":2}
	fmt.Println(myMap2)

	// Checking equality
	if maps.Equal(myMap,myMap2){
	fmt.Println("Both the maps are equal ✅")
	}else{
		fmt.Println("Unequal maps ❌")
	}

	// Iterating over maps
	for k,v:=range myMap3{
		fmt.Println(k,":",v)
	}
	// also: _,v or k,_ ✔️

	myMap4:= map[string]string{"Batman":"Bruce Wayne🦇", "SpiderMan":"Peter Parker🕷️","IronMan":"Tony Stark🚀"}
	fmt.Println(myMap4)

	//! Real life equality checking... OK
	_,ok:=myMap4["SpiderMan"]
	fmt.Println("✅ Is a value associated with SpiderMan:",ok)

	if ok{
		fmt.Println("Value exists with key: SpiderMan ✅")
	}else{
		fmt.Println("Value does not exist with key: SpiderMan ❌")
	}

	// ZERO-VALUE of uninitialized map is: nil

	var myMap5 map[string]string

	if myMap5 == nil{
		fmt.Printf("%v is initialized to nil-value\n",myMap5)
	}else{
		fmt.Printf("%v is initialized, it's not nil\n",myMap5)
	}

	value := myMap5["randomKey"]
	fmt.Println("Non-existent value:",value)

	//myMap5["someKey"] = "someValue"
	//fmt.Println(myMap5) // ❌

	//! 💡 We have to use make()

	myMap5 = make(map[string]string)
	myMap5["someKey"] = "someValue"
	fmt.Println(myMap5) // ✅

	// Length..
	fmt.Println("myMap's length is:",len(myMap))

	// Nested maps / 2d maps
	twoDMap := make(map[string]map[string]string)
	twoDMap["map1"] = myMap4
	fmt.Println(twoDMap)
	// map[map1:map[Batman:Bruce Wayne🦇 IronMan:Tony Stark🚀 SpiderMan:Peter Parker🕷️]]

}