package continuebreak

import "fmt"

func ContinueBreak() {
	for i := 1; i <= 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}
	for i := 1; i <= 5; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
