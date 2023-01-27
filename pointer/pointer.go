package pointer

import "fmt"

func Pointer() {
	var num int = 5
	var name = "Dung"
	var ptr *string
	ptr = &name
	fmt.Println(num)
	fmt.Println(&num)
	fmt.Println(ptr)
	fmt.Println(&name)
	fmt.Println(*ptr)

	name = "Nam"
	fmt.Println(name)
	fmt.Println(*ptr)

	//Change value using pointer
	*ptr = "Mai"
	fmt.Println(name)
	fmt.Println(*ptr)

}
