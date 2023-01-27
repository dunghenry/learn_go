package ie

import "fmt"

type Shape1 interface {
	area() float32
}

type Rectangle1 struct {
	length, breadth float32
}

func (r Rectangle1) area() float32 {
	return r.length * r.breadth
}

type Triangle struct {
	base, heigth float32
}

func (t Triangle) area() float32 {
	return t.base * t.heigth * 0.5
}
func caculate1(s Shape1) {
	fmt.Println(s.area())
}
func MultipleStruct() {
	r := Rectangle1{7, 4}
	t := Triangle{8, 12}
	caculate1(r)
	caculate1(t)
}
