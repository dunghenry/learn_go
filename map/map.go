package mp

import "fmt"

func Map() {
	subjectMarks := map[string]float32{"Go": 4.5, "JS": 1}

	fmt.Println(subjectMarks["Go"])
	// delete(subjectMarks, "JS")

	for name, value := range subjectMarks{
		fmt.Println(name)
		fmt.Println(value)
	}
}
