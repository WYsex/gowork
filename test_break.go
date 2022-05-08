package main

import (
	"fmt"
)

func test2() {
to_do:

	for t := 0; t <= 10; t++ {
		if t == 5 {
			fmt.Println(t)
			break to_do
		}
	}
}

func test1() {
	i := 2
	switch i {
	case 1:
		fmt.Println("不用结束")
	case 2:
		fmt.Println("end。。。")
		break
	case 3:
		fmt.Println("不用结束")
	}

}

func main() {
	test1()
	test2()
}
