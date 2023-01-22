package array

import "fmt"

func MultidimensionalArray() {
	arrInteger := [2][2]int{{1, 2}, {3, 4}}
	for i := 0; i < len(arrInteger[0]); i++ {
		for j := 0; j < len(arrInteger[1]); j++ {
			fmt.Println(arrInteger[i][j])
		}
	}
}
