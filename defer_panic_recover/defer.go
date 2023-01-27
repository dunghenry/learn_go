package deferpanicrecover

import "fmt"

func Defer() {
	defer fmt.Println("3")
	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("----------------------------")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")
}
