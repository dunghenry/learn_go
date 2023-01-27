package ie

import "fmt"

type Shape interface {
	area() float32
	perimeter() float32
}

type Rectangle struct {
	length, breadth float32
}

// use struct to implement area() of interface
func (r Rectangle) area() float32 {
	return r.length * r.breadth
}

// use struct to implement area() of interface
func (r Rectangle) perimeter() float32 {
	return (r.length + r.breadth) * 2
}
func caculate(s Shape) {
	fmt.Println(s.area())
}
func Interface() {
	rect := Rectangle{2, 3}
	caculate(rect)
	
}
