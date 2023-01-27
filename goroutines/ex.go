package goroutines

import (
	"fmt"
	"time"
)

func dis(msg string) {
	for {
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
func Ex() {
	go dis("Xin chao")
	dis("Hhh")
}
