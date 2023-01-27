package deferpanicrecover

import "fmt"
func division(a, b int){
	if b == 0{
		panic("Cannot divide a number by zero")
	}else{
		rs := a / b
		fmt.Println(rs)
	}
}
func Panic() {
	fmt.Println("Help! Something bad is happening")
	// panic("Ending the program")
	fmt.Println("Waiting to execute")
	division(1, 2)
	division(8, 0)
	division(1, 3)
}
