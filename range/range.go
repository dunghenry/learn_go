package re

import "fmt"

func Range() {
	numbers := [5]int{1, 2, 3, 4, 5}

	//array
	for index, item := range numbers {
		fmt.Print("Index")
		fmt.Println(index)
		fmt.Print("Item")
		fmt.Println(item)
	}

	// string
	str := "Hello"
	for index, item := range str {

		fmt.Println(index)

		fmt.Printf("%c", item)
		println()
	}

	//map
	subjectMarks := map[string]int{"Java": 5, "Golang": 6}
	for subject, mark := range subjectMarks {
		fmt.Println(subject, " : ", mark)
	}
}
