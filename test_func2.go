package main

import "fmt"

func sunm(a int, b int) string {
	sunb := a + b
	return string(sunb)
}

func max(a int, b int) string {
	if a > b {
		return "a大"
	} else if b < a {
		return "b大"
	} else if a == b {
		return "一样大"
	}
	return "null"
}
func main() {
	type ff func(a int, b int) string
	var g ff
	g = max
	g = sunm
	h := g(1, 1)
	fmt.Println(h)
}
