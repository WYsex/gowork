package main

import (
	"fmt"
)

func test(name string, age string) string {
	a := name
	b := age
	f1 := func() string {
		return a + b
	}
	msg := f1()
	return msg
}

func main() {
	//匿名函数
	max := func(a, b int) int {
		if a < b {
			return b
		} else {
			return a
		}
	}
	fmt.Println(max(1, 3))
	//匿名函数自己调用自己
	zijidiaoyong := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}(1, 2)
	fmt.Println(zijidiaoyong)
	f1 := test("zhangsan", "20")
	fmt.Println(f1)
}
