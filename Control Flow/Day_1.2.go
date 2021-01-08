package main

import "fmt"

var v1 int = 400
var v2 int = 700

if v1 == 400 {
	if v2 == 700 {
		fmt.Printf("Value of v1 is 400 and v2 is 700")
	}
}

// ----------- if-else-if condition
if v1 == 100 {
	fmt.Printf("Value of v1 is 100\n")
}else if v1 == 200{
	fmt.Printf("Value of a is 20\n")
}else if v1 ==300 {
	fmt.Printf("Value of a is 300\n")
}else{
	fmt.Printf("None of the varialble is matching")
}