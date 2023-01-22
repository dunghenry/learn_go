package operators

import "fmt"

func ArithmeticOperator() {

	num1 := 5
	num2 := 2

	num3 := 4.0
	num4 := 1.5
	//addition
	fmt.Printf("%d + %d = %d\n", num1, num2, num1+num2)
	//subtraction
	fmt.Printf("%d - %d = %d\n", num1, num2, num1-num2)
	//multiplication
	fmt.Printf("%d * %d = %d\n", num1, num2, num1*num2)
	//division
	fmt.Printf("%g / %g = %g\n", num3, num4, num3/num4)
	//modulo division
	fmt.Printf("%d %% %d = %d\n", num1, num2, num1%num2)
}
