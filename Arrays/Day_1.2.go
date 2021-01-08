package main

import "fmt"

func main(){
	arr := [7]string{"one","two","three","four"}

	fmt.Println("Arrays :",arr)

	//creating a slice
	myslice := arr[1:6]

	fmt.Println("Slice: ",myslice)

	//display length of the slice
	fmt.Printf("Length of the slice : %d",len(myslice))

	//display the capacity of slice
	fmt.Printf("\nCapacity of the slice :%d",cap(myslice))


}