package main

import "fmt"

func main(){
	str1:="hello"
	str2:="world"
	result1:=str1==str2

	//display the result
	fmt.Println(result1)

	//display type of result1
	fmt.Println("The type of result is : %T\n",result1)

	str:="gfg"

	//display length of string
	fmt.Printf("\nString is :%s",str)

	//display the type of string 
	fmt.Printf("\n Type of Str is :%T",str)
}