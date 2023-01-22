package st

import "fmt"

type User struct {
	name string
	age  int
}

func Struct() {

	var user User
	user1 := User{
		name: "Nam",
		age:  21}
	user2 := User{"Mai", 11}
	user = User{
		name: "Dung",
		age:  21,
	}
	fmt.Println(user)
	fmt.Println(user1)
	fmt.Println(user2)
}
