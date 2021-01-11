 //advanyages
package main 
  
import ( 
    "fmt"
    "time"
) 

func main(){
    fmt.Println("Welcome to main function");

    go func(){
        fmt.Println("welcome o gfg");
    }()

    time.Sleep(1 * time.Second)
    fmt.Println("Good bye to main function")


}