package pointerstruct

import "fmt"

type Person struct {
	name string
	age  int
}

func PointerStruct() {
	person1 := Person{"Dung", 21}
	var ptr *Person
	ptr = &person1
	fmt.Println(ptr.name)
	fmt.Println(ptr.age)
	fmt.Println(ptr)
	fmt.Println((*ptr).name)
	//Change value
	ptr.age = 25


}
