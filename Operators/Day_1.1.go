package main

import "fmt"

func main(){
	p:=23
	q:=60

	//Arithmatic Operator
	result1:=p+q
	fmt.Printf("Rsult of p+ q =%d\n",result1)

	//Relational Operators
	result2:=p==q
	fmt.Println(result2)

	//relational Operaters !=
	result3:=p != q 
	fmt.Println(result3)

	if p != q || p<=q{
		fmt.Println("True")
	}

	//bitwise Operater
	result4 := p & q
	fmt.Printf("Result of P & Q =%d\n",result4)

	//Assignment Operaors '='
	p=q 
	fmt.Println(p)
}