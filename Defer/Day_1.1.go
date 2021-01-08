package main

import "fmt"

func mul(a1,a2 int)int{
	res := a1 * a2
	fmt.Println("Result is :",res)
	return 0
}

func show(){
	fmt.Println("Hello ,world")
}

//main function
func main(){
	mul(22,45)

	defer mul(23,56)

	show()
}