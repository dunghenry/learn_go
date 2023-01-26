package anonymousfunction

import "fmt"

func AnonymousFunction() {
	var hello = func() {
		fmt.Println("Hi")
	}
	hello()
}
