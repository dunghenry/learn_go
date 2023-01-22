package array

import "fmt"

func Array() {
	// var n int
	// var arrInt = [4]int{1, 2, 3, 4}
	// languages = [3]string{"JS", "C", "C++"}
	var arrCourse = [4]string{0: "JS", 3: "React"} // ["JS", "", "", "React"]

	// var age = [...]int{1, 2, 3, 4}
	// age := [...]int{1, 2, 3, 4}
	age := []int{1, 2, 3, 4}

	// fmt.Scan(&n)
	// var array []int = make([]int, int(n))
	// for i := 0; i < len(array); i++ {
	// 	fmt.Scan(&array[i])
	// }

	// for i := 0; i < len(array); i++ {
	// 	fmt.Println(array[i])
	// }

	for index, value := range arrCourse {
		fmt.Println(index)
		fmt.Printf("%s", value)
		fmt.Println()
	}
	fmt.Println(len(age)) //4
}
