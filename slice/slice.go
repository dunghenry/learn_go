package slice

import "fmt"

func Slice() {
	arr := [5]int{1, 2, 3, 4, 5}
	// nums := []int{1, 2, 3, 4, 5}
	// var numbers = []int{1, 2, 3, 4}
	numbers := arr[1:5]
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	//Slice function: append(), copy(), len(), Equal()
	slice := make([]string, 3, 5)
	slice = append(slice, "JS")
	slice = append(slice, "Java")
	slice = append(slice, "C")
	slice = append(slice, "C++")
	slice = append(slice, "Ruby")
	slice = append(slice, "React")
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
}
