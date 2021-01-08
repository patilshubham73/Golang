package main

import "fmt"

type tank interface {
	//methods
	Tarea() float64
	Volume() float64
}

type myvalue struct{
	radius float64
	height float
}

//implementingmethods of tank interfaces
func (m myvalue) Tarea() float64{
	return 2*3.radius*m.height+2*3.14*m.radius*m.radius
}

func (m myvalue) Volume() float64{
	return 3.14* m.radius * m.radius * h.height
}

func main(){

	var t tank
	t = myvalue{10,14}
	fmt.Println("Area of thank :",t.Tarea())
	fmt.Println("Voulume Of Tank :",t.Volume())
}