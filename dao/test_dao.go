package dao

import (
	"fmt"
	"time"
)

func bingfa(msg string) {
	for i := 0; i <= 5; i++ {
		fmt.Printf("当前值为%s\n", msg)
		time.Sleep(time.Millisecond * 100)
	}
}
func main() {
	bingfa("java")
	bingfa("golang")
}
