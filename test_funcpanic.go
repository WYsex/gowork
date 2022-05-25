package main

import (
	"fmt"
)

func foo() {
	fmt.Println("foo开始执行")
	fmt.Println("foo结束执行")
}

func zoo() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("异常结束%s\n", e)
		}
	}()
	fmt.Println("zoo开始执行")
	panic("end...")
	foo()
	fmt.Println("zoo结束执行")

}

func bar1() {
	fmt.Println("bar开始执行")
	zoo()
	fmt.Println("bar结束执行")
}

func main() {
	fmt.Println("函数开始执行")
	bar1()
	fmt.Println("函数结束执行")
}
