package recursion

import "fmt"

func sum(num int) int {
	if num == 0 {
		return 0
	} else {
		return num + sum(num-1)
	}
}
func factorial(num int) int {
	if num == 0 {
		return 1
	} else {
		return num * factorial(num-1)
	}
}
func Recursion() {
	fmt.Println(sum(10))
	fmt.Println(factorial(3))
}
