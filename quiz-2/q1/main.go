package main

import (
    "fmt"
)
// create interface 
type shape interface {
    area() float64
    perimeter() float64
}
	
type rectangle struct {
    width, height float64
}

func (r rectangle) area() float64 {
    return r.width * r.height
}
func (r rectangle) perimeter() float64 {
    return 2*r.width + 2*r.height
}
func measure(s shape) {
    fmt.Println("your area is: ",s.area())
    fmt.Println("your perimeter is: ",s.perimeter())
}

func main() {
    r := rectangle{width: 3, height: 4}
	measure(r)
}
