package main

import (
	"fmt"
)

type Position struct {
	x float64
	y float64
}
type Daxiao struct {
	x int
	t int
}

func main() {
	//map的声明及初始化
	// mp := map[int]int{
	mp := map[Position]string{
		Position{29.935523, 52.568915}:  "school",
		Position{25.352594, 113.304361}: "shopping-mall",
		Position{73.224455, 111.804306}: "hospital"}
	fmt.Println(mp)

	mp1 := map[Daxiao]string{
		{1, 3}: "前者小",
	}
	fmt.Println(mp1)

	//使用make对map进行显式初始化
	mp2 := make(map[int]string)    //未指定初始容量
	mp3 := make(map[int]string, 9) //指定初始容量为9
	fmt.Println(mp2)
	fmt.Println(mp3)
	mp3[10] = "第9⃣️"
	fmt.Println(mp3)
	mp4 := map[string]int{
		"key1": 1,
		"key2": 2,
	}
	fmt.Println(mp4)
	mp4["key1"] = 11 //key1的值会覆盖
	fmt.Println(mp4)
	fmt.Println("map的长度为%d", len(mp4))
	//fmt.Println(cap(mp4))  //map不支持cap方法
	_, ok := mp4["key4"]
	if !ok {
		fmt.Println("当前key不存在")
	} else {
		fmt.Println(ok)
	}
	delete(mp4, "key2") //删除key2
	fmt.Println(mp4)
	fmt.Println("--------------------")
	for k, v := range mp4 {
		fmt.Printf("{[%s,%d]}\n", k, v)
	}
	for k := range mp4 {
		fmt.Printf("{%s}", k)
	}
	//不要依赖map的元素遍历顺序
	//map不是线程安全的，不支持并发读写
	//不要尝试获取map中value的地址
}
