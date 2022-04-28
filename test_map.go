package main

import "fmt"

type Position struct {
	x float64
	y float64
}

func main() {
	//map的声明及初始化
	// mp := map[int]int{
	mp := map[Position]string{
		Position{29.935523, 52.568915}:  "school",
		Position{25.352594, 113.304361}: "shopping-mall",
		Position{73.224455, 111.804306}: "hospital"}
	fmt.Println(mp)
}
