package sw

import "fmt"

func Switch() {
	println()
	s := 2
	switch {
	case s <= 1:
		fmt.Println(1)
		fallthrough
	case s <= 2:
		fmt.Println(1, 2)
		fallthrough
	case s <= 3:
		fmt.Println(1, 2, 3)
	}

	dayOfWeek := "Monday"
	switch dayOfWeek {
	case "Saturday", "Sunday":
		fmt.Println("Weekend")
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Weekday")
	default:
		fmt.Println("Invalid day")
	}
}
