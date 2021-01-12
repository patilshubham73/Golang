package main

import (
	"fmt"
	"time"

)


func main(){

	fmt.Println("Welcome to main Function")
	
	go func(){
		fmt.Println("Welcome to golang")
	}()
	
	time.Sleep(1* time.Second)
	fmt.Println("bye bye")
}