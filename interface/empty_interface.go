package ie

import "fmt"

func displayValue(i interface{}) {
	fmt.Println(i)
}
func display(i ...interface{}) {
	fmt.Println(i)
}
func EmptyInterface() {
	var a interface{}
	fmt.Println(a) //nill
	b := "Hello"
	c := 20
	d := false
	displayValue(b)
	displayValue(c)
	displayValue(d)
	display(1, 2, 3, 4, 5, 6, 7)
}
