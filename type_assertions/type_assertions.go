package typeassertions

import "fmt"

func TypeAssertions() {
	var a interface{}
	a = 12
	interfaceValue, booleanValue := a.(int)
	fmt.Println(interfaceValue)
	fmt.Println(booleanValue)

}
