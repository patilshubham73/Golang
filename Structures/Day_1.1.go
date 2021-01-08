package main  

import "fmt"

type Address struct {

	Name string
	city string
	pincode int
}

func main(){

	var a Address
	fmt.Println(a)

	a1 := Address{"Akshay","patil",5685632563}

	fmt.Println("Address : ",a1)

	a2 := Address{Name: "ankita",city:"pune",pincode:56366}

	fmt.Println(Address :,a2)

	a3:= Address{Name :"Delhi"}
	fmt.Println("address ",a3)


}