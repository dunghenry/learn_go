package goroutines

import "fmt"

func display(msg string) {
	fmt.Println(msg)
}
func Goroutines() {
	go display("Hehe")
	display("Hi")
}
