package main

import (
	"fmt"
	"unsafe"
)

func foo(arr [5]int) {

}
func main() {
	var arr = [6]int{1, 2, 3, 4, 5, 6}
	var arr1 [5]int    //没有初始化，int型值为0
	var arr2 [5]string //没有初始化，int型值为0

	fmt.Println("数组长度：", len(arr))
	fmt.Println("数组大小：", unsafe.Sizeof(arr))
	fmt.Println(arr1)
	fmt.Println(len(arr2))

}
