package ifelse

import "fmt"

func IfElse() {
	n := 1
	m := 2
	println()
	if n > m {
		println("%d > %d", n, m)
	} else {
		fmt.Printf("%d < %d", n, m)
	}
	//else if (condition)
}
