package typecasting

import "fmt"

func TypeCasting() {
	println()
	var floatValue float32 = 9.8
	var intValue int = 1
	var convertValue int = int(floatValue)
	println(convertValue) //9
	println(intValue)
	kq := float32(intValue)
	fmt.Printf("%.2f", kq) //
}
