package takeinput

import "fmt"

func InputData() {
	// fmt.Scanln(), fmt.Scanf(), fmt.Scan()
	fmt.Scan()
	var a int
	fmt.Printf("Nhap a: ")
	fmt.Scan(&a)
	var msg string
	fmt.Printf("Nhap msg: ")
	fmt.Scanf("%s", &msg)
	fmt.Print(a)
	fmt.Println()
	fmt.Printf("%T", a)
	println()
}
