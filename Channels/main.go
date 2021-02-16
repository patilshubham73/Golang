package  main

import {
	"fmt"
	"time"
}

func main() {
	c := make(chan int)

	go func(i:=0;i<10;i++){
		c <-i
	}()

	go func(){
		for{
			fmt.Println(<-c)
		}
	}()

	time.sleep(time.Second)

}