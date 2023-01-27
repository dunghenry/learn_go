package pointerfunction

import "fmt"

func update(num *int) {
	*num = 10
}

//return pointer from function

func display() *string {
	msg := "Hello"
	return &msg
}

// call by value
func callByValue(num int) {
	num = 30
	fmt.Println(num)
}

//call by reference
func callByReference(num *int) {
	*num = 1
	fmt.Println(*num)
}

func PointerFunction() {
	var number = 100
	update(&number)
	fmt.Println(number)
	rs := display()
	fmt.Println(*rs)
	var num int = 50
	callByValue(num)
	fmt.Println(num)
	callByReference(&num)
	fmt.Println(num)
}
