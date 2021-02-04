package main

import (
	
	"math"
) 

type circle struct {
	redius float64
}

type rect struct {
	width float64
	height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}


func main(){

	c1 := circle(12.5)
	r1 := rect(8,5)


}