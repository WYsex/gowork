package main

import "fmt"

func myAppend(sl []int, elems ...int) []int {
	fmt.Printf("%T\n", elems)
	if len(elems) == 0 {
		fmt.Println("没有往elems append内容")
		return sl
	}
	sl = append(sl, elems...)
	fmt.Println("111", sl)
	return sl
}

func teup(task string) func() {
	fmt.Printf("已进入外部函数体%s\n", task)
	return func() {
		fmt.Printf("此名称为%s\n", task)
	}
}

func main() {
	teardown := teup("你好")
	defer teardown()
	fmt.Println("nihao")
	sl := []int{1, 2, 3}
	sl = myAppend(sl)
	fmt.Printf("内容为%d\n", sl)
	sl = myAppend(sl, 4, 5, 6)
	fmt.Printf("内容2为%d\n", sl)
}
