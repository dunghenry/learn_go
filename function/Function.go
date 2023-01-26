package function

import "fmt"

func getMsg() string {
	return "Hello, world!"
}
func sayHello() {
	fmt.Println("Hello")
}
func sum(a int, b int) int {
	a += 1
	b += 1
	return a + b
}
func Function() {
	a, b := 1, 2
	fmt.Println(getMsg())
	sayHello()
	fmt.Println(sum(a, b))
	fmt.Println(a, b)
}
