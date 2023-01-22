package st

import "fmt"

type Rectangle func(int, int) int

type RectanglePara struct {
	length  int
	breadth int
	color   string
	rect    Rectangle
}

func FunctionInsideStruct() {
	rs := RectanglePara{
		length:  10,
		breadth: 20,
		color:   "Red",
		rect: func(length int, breadth int) int {
			return length * breadth
		},
	}
	fmt.Println(rs.color)
	fmt.Println(rs.rect(rs.length, rs.breadth))
}
