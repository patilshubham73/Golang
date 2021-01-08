package main

import "fmt"

func main(){

	//using 8 bit unsigned int
	var x uint8=225
	fmt.Println(x+1,x)

	//using 16 bit signed int
	var y int16=32767
	fmt.Println(y+2,y-2)

	a:=20.45
	b:=53.56

	var m complex128=complex(6,2)
	var n complex64=complex(9,2)
	
	c:=b-a

	//display the result
	fmt.Println("Result is :%f\n",c)

	//display the type of c variable
	fmt.Printf("the type is c: %T\n",c)

	fmt.Println(m)
	fmt.Println(n)

	//display the type
	fmt.Print("the type of m is %T and type of n is %T",m,n)


}