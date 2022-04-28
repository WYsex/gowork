package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var arr = [6]int{1, 2, 3, 4, 5, 6}
	var arr1 [5]int    //没有初始化，int型值为0
	var arr2 [5]string //没有初始化，int型值为0
	var arr3 = [...]string{"1", "2", "3"}
	var arr4 = [...]int{10: 12} //将下标为10的值赋值12 其余值为0

	var arr5 = []int{1, 2, 3, 4, 5, 6, 7, 8} //初始化slice
	arr5 = append(arr5, 9)

	//创建切片，初始长度为1，最大长度为7
	arr6 := make([]int, 2, 7)
	for i := 0; i <= 7; i++ {
		arr6 = append(arr6, i)

	}
	fmt.Println("数组长度：", len(arr))
	fmt.Println("数组大小：", unsafe.Sizeof(arr))
	fmt.Println(arr1)
	fmt.Println(len(arr2))
	fmt.Println(len(arr3))
	fmt.Println(arr4)
	fmt.Println(arr5)
	fmt.Println(arr6)
	fmt.Println(len(arr6)) //9
	fmt.Println(cap(arr6)) //14

}
