package variablescope

import "fmt"

var sum int = 1 //global
const PI = 3.14

func addNumbers() int {
	sum := 1 + 2 // local
	return sum
}
func VariableScope() {
	fmt.Println(sum)
	fmt.Println(PI)
	fmt.Println(addNumbers())
}