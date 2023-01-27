package deferpanicrecover

import "fmt"

func handlePanic() {
	a := recover()
	if a != nil {
		fmt.Println("Recover", a)
	}
}
func division1(a, b int) {
	defer handlePanic()
	if b == 0 {
		panic("Cannot divide a number by zero")
	} else {
		rs := a / b
		fmt.Println(rs)
	}
}
func Recover() {
	division1(4, 2)
	division1(8, 0)
	division1(2, 8)
}
