package string

import (
	"fmt"
	"strings"
)

func String() {
	var msg string = "Hello"
	message := `Xin chao
	
	hihi.Xin chao
	`
	fmt.Println(msg)
	fmt.Println(message)
	// len(message)
	// string method: Compare()
	fmt.Println(strings.Compare(msg, message))      // -1: msg < message
	fmt.Println(strings.Contains(message, "Xin"))   // true
	fmt.Println(strings.Contains(message, "Hello")) // false
	str := strings.Replace(msg, "l", "L", 2)
	fmt.Println(str) //HeLLo
	// ToUpper(), ToLower()

	//split
	fullName := "Tran Van Dung"
	fmt.Println(strings.Split(fullName, " "))
	arrStr := []string{"Tran", "Van", "Dung"}
	fmt.Println(strings.Join(arrStr, " "))
}
