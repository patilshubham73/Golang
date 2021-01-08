package main

import "fmt"

func main(){
	var myarr []string

	myarr[0]="one"
	myarr[1]="two"

	fmt.Println("Elements of Arrays :")
	fmt.Println("Elements 1 : ",myarr[0])
	fmt.Println("Elements 2 : ",myarr[1])

	arr:= [4]string{"one","two","three","fourth"}

	fmt.Println("\nPrintln of the arrays :")

	for i:=0; i<3;i++{
		fmt.Println(arr[i])
	}
}